package connection

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Silent 静默 不输出任何日志
		// Error 、Warn、Info(输出所有 SQL 语句)
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("mysql 连接失败 error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(&User{}, &Comment{}, &Post{})
	if err != nil {
		return nil, fmt.Errorf("根据实体类生成表结构失败: %v", err)
	}

	return db, nil
}
