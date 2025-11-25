package logic

import (
	"context"

	"blog-system-zero/app/blog/internal/svc"
	"blog-system-zero/app/blog/types/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *blog.DeleteCommentRequest) (*blog.BaseResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.BaseResponse{}, nil
}
