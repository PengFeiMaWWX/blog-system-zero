package logic

import (
	"context"

	"blog-system-zero/app/blog/internal/svc"
	"blog-system-zero/app/blog/types/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostRequest) (*blog.BaseResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.BaseResponse{}, nil
}
