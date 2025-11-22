package logic

import (
	"context"

	"blog/internal/svc"
	"blog/types/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsLogic) GetComments(in *blog.GetCommentsRequest) (*blog.GetCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.GetCommentsResponse{}, nil
}
