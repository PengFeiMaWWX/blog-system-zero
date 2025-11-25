package main

import (
	"blog-system-zero/common/model"
	"blog-system-zero/etcd"
	"flag"
	"fmt"
	"log"

	"blog-system-zero/app/user/internal/config"
	"blog-system-zero/app/user/internal/server"
	"blog-system-zero/app/user/internal/svc"
	"blog-system-zero/app/user/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	log.Printf("connect to etcd success at %s...\n", c.Etcd.Hosts)
	manager, err := etcd.NewConfigManager(c.Etcd.Hosts)
	if err != nil {
		log.Fatalf("创建etcd管理器失败： %v", err)
	}

	// 使用 etcd 获取数据库配置
	dbConfigKey := "/config/database" // 你的数据库配置在 etcd 中的 key
	dsn, err := manager.GetDbConfig(dbConfigKey, "")
	if err != nil {
		log.Fatalf("获取数据库配置失败 ： %v", err)
	}

	log.Printf("成功获取数据库 DSN: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 创建 serviceContext 时传入 db
	ctx := svc.NewServiceContext(c, db)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
