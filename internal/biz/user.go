package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

// 模型
type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
}

type UserRepo interface {
	// 创建用户
	Create(context.Context, *User) (*User, error)
	// Update(context.Context, *User) (*User, error)

	// 根据用户名查找用户
	FindByUsername(context.Context, string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// Login 登录
func (uc *UserUsecase) Login(ctx context.Context, u *User) (*User, error) {
	// 1. 查找用户是否存在
	user, err := uc.repo.FindByUsername(ctx, u.Username)
	if err != nil {
		return nil, err
	}
	// 2. 比对密码是否正确
	oldPassword := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	if strings.Compare(user.Password, oldPassword) != 0 {
		return nil, errors.New("密码错误")
	}
	return &User{ID: user.ID, Username: user.Username}, nil
}

// Register 注册
func (uc *UserUsecase) Register(ctx context.Context, u *User) (*User, error) {
	// 1. 检查用户是否存在
	_, err := uc.repo.FindByUsername(ctx, u.Username)
	if err == nil {
		// 用户已存在
		return nil, errors.New("用户已存在")
	}
	// 2. 加密密码
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	// 3. 入库
	user, err := uc.repo.Create(ctx, &User{Username: u.Username, Password: pwd})
	if err != nil {
		return nil, err
	}
	return user, nil
}
