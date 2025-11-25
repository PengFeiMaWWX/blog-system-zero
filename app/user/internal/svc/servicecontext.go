package svc

import (
	"blog-system-zero/app/user/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB // 添加 DB 字段

}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
