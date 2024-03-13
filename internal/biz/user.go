package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
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
	// 根据用户ID查找用户
	FindById(context.Context, uint) (*User, error)
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

// GenerateToken 生成Token
func (uc *UserUsecase) GenerateToken(ctx context.Context, u *User) (string, error) {
	privateKeyFile, err := os.Open("./config/private.key")
	if err != nil {
		return "", err
	}
	privateKeyByte, err := io.ReadAll(privateKeyFile)
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"iss":      "my-auth-server",
		"uid":      u.ID,
		"username": u.Username,
	})
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return "", err
	}
	s, err := t.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return s, nil
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

func (uc *UserUsecase) ValidToken(ctx context.Context, u *User) (bool, *User, error) {
	u, err := uc.repo.FindById(ctx, u.ID)
	if err != nil {
		return false, nil, err
	}
	return true, u, nil
}
