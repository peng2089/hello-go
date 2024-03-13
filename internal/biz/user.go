package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
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
	key := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCYCiJxzjXHSpNg
c0aX73J3MXpU3NTzYXkVbkfU8kMYvsKMqShKD+aMme0bcX8fjLqh5zbDMHBvsclx
IWFvWf7w++WTAZWVSefmErzwoetU5+EQmwpk4qn0FpC1HKOOZ7TdwQTyJCBnr2QX
pQ4NaRBZe8QMJ301aHixAiqmj+fHmmWMHue5l7kZXd1JTUdNDWwTMCRMI+7uWeu0
bNJ7wRrgE8XLrYlmXgxIz2I6YeyOYN74oiUvWrqVAhISpJIQn9jHwYRcOJsnvE8Y
Hhg689wqjGubHxcX2VKPwMVt9NAQTiy9hOKdb5LJ64VlAw/TmD2E5bCqr/fWvMQA
Os3M6BAFAgMBAAECggEAGQEBrSm2mnHfTutmXrJYZtXSQoaW1vfey/F5CsJU2or0
v+FJ6PQ4fEMMRYki2PNx9hJqZStgMl5QvLQ6q+9nCAbFOKn4Cbc/1gya2hAm/a2s
y+hTi0fjn2renYp6M39GtXl68L+UPLkRgvn4F6iBsdWy5jTQWKo3vxSWIxJjkeXH
jVa/KpEj5G+lpTWmuckKDi5PAFHeIjymHu5y0Rgz0PJE8ntg1pNFvySwstcD4iNP
Q3WYpse4wBuM2ocx6jx9aQ+OjyLzHpc/TlethJ8swKUyCXSqI0SX1Uiz89ZWFdF5
sym3XqYaDe31oxpnaUW1tbh2bdTcf5ytRX4jlCRSgQKBgQDFeHkcGbSz8+U8Xoc+
pV13Gp6YrKKgExb+DWU4TSzRdQQi51faQre/8HF2wPeh3U/agm68paU7lCa8bpWh
gHJ0rWM4D2xAo8TrkxXRC/K2sSWE1x0ewcq1ypcGTQsC3kFUrWmb+3LK8+tp5mLb
hWz7HnaY03UEdg4PdHHhYtj+0QKBgQDFGnssEweWYG9Q2invxraQVAFjwIwg0oyx
nf2JkPvTArE/fs+AQPUBR1mZxzvc7fxqMojRdziSPyCq9NyfTAv02+CcbE4g/7S+
KtiWUYYHydPoRTPPxOoaX8bEGop2ZtZ2kbY4zOVweZq+aKFhNFj7iZU8RPtVszd/
DWPnd5mS9QKBgHpB10GVjLIpG1Df+X3IpfA6k9xUba8Lgp2xr2xHI6tedjwh3Ntt
mRQFikoKuIYHXgwb2vGC4KTdWYoqMQu3WaVxP8+ShOQUQLPU8ZnmetOvI1p2UDod
oyIgFpa5FkslTW9emdcPu4d5stNy5tugZNOojaaarIUhjhz1bBgtuckxAoGAKxUn
uTGLpgX8LusQ4ZVI0HFcQGaU+pOrJyPGiGxFbxEWly9rwbfxFc93uVJANoFafAgB
ue9aUTU1Ocj99F/V+TaaePZ/eV0LL8oDv4+gQVGPXSTLN23uUcd/ldvLSigeVppw
/ydiO1yJQ3dxVuLvVEP1d9AIM+pRhhbyMGUHD4UCgYEAszaS+/EUUtRzXv/TilP+
Yd6A1xtyOFUNFmbxPpcgZSfa9iyZt6yt/8VFrAUliOSR2tzwQJMaIpKfWz0UmnNz
trEOOuqJWQ5A5zRE0QsP10ulAf2zgm+e/Qh5VyxYCOZ0bn40BrY7WO3QIuyGaN5h
4LCzHxhhACpD8lFBmbjbNSQ=
-----END PRIVATE KEY-----`
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"iss":      "my-auth-server",
		"uid":      u.ID,
		"username": u.Username,
	})
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
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

func (uc *UserUsecase) ValidToken(ctx context.Context, u *User) (bool, error) {
	_, err := uc.repo.FindById(ctx, u.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
