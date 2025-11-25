package logic

import (
	"blog-system-zero/common/auth"
	"blog-system-zero/common/model"
	"context"
	"errors"

	"blog-system-zero/app/user/internal/svc"
	"blog-system-zero/app/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	var userResult model.User
	err := l.svcCtx.DB.Where("username = ?", in.Username).First(&userResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户存在！")
		}
		return nil, err
	}

	//err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(in.Password))
	//if err != nil {
	//	return nil, errors.New("用户名或密码错误")
	//}

	token, err := auth.GenerateJWT(userResult.ID, userResult.Password, 10000)
	if err != nil {
		l.Logger.Errorf("Token生成失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 手动映射到 protobuf 生成的 User 类型
	userProto := &user.UserInfo{
		Id:       userResult.ID,
		Username: userResult.Username,
		Email:    userResult.Email,
	}
	return &user.LoginResponse{
		Token: token,
		User:  userProto,
	}, nil
}
