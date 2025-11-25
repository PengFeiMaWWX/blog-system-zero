package logic

import (
	"context"

	"blog-system-zero/app/blog/internal/svc"
	"blog-system-zero/app/blog/types/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *blog.CreateCommentRequest) (*blog.BaseResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.BaseResponse{}, nil
}
