package logic

import (
	"blog-system-zero/app/user/internal/svc"
	"blog-system-zero/app/user/types/user"
	"blog-system-zero/common/model"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	// 1 参数验证
	if err := l.validateRequest(in); err != nil {
		return l.errorResponse(1001, "参数验证失败: "+err.Error())
	}
	// 2. 从数据库查询用户信息
	userInfo, err := l.getUserFromDB(in.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return l.errorResponse(1002, "用户不存在！")
		}
		l.Logger.Error("查询用户失败， userId：%d, error: %v", in.UserId, err)
		return l.errorResponse(500, "系统错误")
	}

	// 3. 构建成功响应
	return l.buildSuccessResponse(userInfo)
}

// validateRequest 验证请求参数
func (l *GetUserLogic) validateRequest(in *user.GetUserRequest) error {

	if in == nil || in.UserId < 0 {
		return errors.New("用户ID必须大于0")
	}
	return nil
}

// buildSuccessResponse 构建成功响应
func (l *GetUserLogic) buildSuccessResponse(userInfo *model.User) (*user.GetUserResponse, error) {

	return &user.GetUserResponse{
		Base: &user.BaseResponse{
			Code:    0,
			Message: "成功",
		},
		User: &user.UserInfo{
			Id:       userInfo.ID,
			Username: userInfo.Username,
		},
	}, nil
}

func (l *GetUserLogic) errorResponse(i int, s string) (*user.GetUserResponse, error) {
	return &user.GetUserResponse{
		Base: &user.BaseResponse{
			Code:    -1,
			Message: "参数验证失败",
		},
	}, nil
}

func (l *GetUserLogic) getUserFromDB(id uint64) (*model.User, error) {

	var userInfo model.User

	// 使用gorm查询用户信息
	result := l.svcCtx.DB.Where("id = ?", id).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userInfo, nil
}
