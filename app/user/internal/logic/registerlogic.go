package logic

import (
	"blog-system-zero/common/auth"
	"blog-system-zero/common/model"
	"context"
	"errors"

	"blog-system-zero/app/user/internal/svc"
	"blog-system-zero/app/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	// 参数验证
	if err := l.validateRequest(in); err != nil {
		return nil, err
	}

	// 检查用户是否存在

	return &user.RegisterResponse{}, nil
}

// checkUserExists 检查用户是否已存在
func (l *RegisterLogic) checkUserExists(in *user.RegisterRequest) error {
	// 检查用户名是否已存在
	var userResult model.User
	err := l.svcCtx.DB.Where("username = ? OR email = ?", in.Username, in.Email).First(userResult).Error
	if err == nil {
		return errors.New("用户已存在！")
	}

	password, err := auth.HashPassword(in.Password)
	if err != nil {
		return errors.New("不能生成hash密码")
	}

	// 创建用户
	user := model.User{
		Username: in.Username,
		Email:    in.Email,
		Password: password,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return errors.New("创建用户失败！")
	}
	return nil
}

// validateRequest 验证请求参数
func (l *RegisterLogic) validateRequest(in *user.RegisterRequest) error {
	if in.Username == "" {
		return errors.New("用户名不能为空")
	}
	if len(in.Username) < 3 || len(in.Username) > 20 {
		return errors.New("用户名长度必须在3-20个字符之间")
	}
	if in.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if in.Password == "" {
		return errors.New("密码不能为空")
	}
	if len(in.Password) < 6 {
		return errors.New("密码长度不能少于6位")
	}

	// 简单的邮箱格式验证
	if !l.isValidEmail(in.Email) {
		return errors.New("邮箱格式不正确")
	}

	return nil
}

// isValidEmail 简单的邮箱格式验证
func (l *RegisterLogic) isValidEmail(email string) bool {
	// 这里可以实现更复杂的邮箱验证逻辑
	return len(email) > 3 && len(email) < 254
}
