package logic

import (
	"context"

	"blog/internal/svc"
	"blog/types/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostsLogic) GetPosts(in *blog.GetPostsRequest) (*blog.GetPostsResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.GetPostsResponse{}, nil
}
