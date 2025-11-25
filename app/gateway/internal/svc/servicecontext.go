// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"gateway/internal/config"
	"gateway/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
