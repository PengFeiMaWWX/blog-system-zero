package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	//Etcd struct {
	//	Hosts []string
	//	Key   string
	//}
}
