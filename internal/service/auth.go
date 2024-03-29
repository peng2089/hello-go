package service

import (
	"hello-go/internal/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	uc *biz.UserUsecase
}

func NewAuthService(uc *biz.UserUsecase) *AuthService {
	return &AuthService{uc: uc}
}

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20" example:"test"`
	Password string `form:"password" json:"password" binding:"required,min=6" example:"123456"`
}

type LoginResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Login: 登录
func (s *AuthService) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	// fmt.Printf("req: %+v\n", req)
	u, err := s.uc.Login(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "用户不存在",
		})
		return
	}

	token, err := s.uc.GenerateToken(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "token生成失败" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": &LoginResp{
			ID:       u.ID,
			Username: u.Username,
			Token:    token,
		},
	})
}

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20" example:"test"`
	Password string `form:"password" json:"password" binding:"required,min=6" example:"123456"`
}
type RegisterResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

// Register: 注册
func (s *AuthService) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	u, err := s.uc.Register(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": &RegisterResp{
		ID:       u.ID,
		Username: u.Username,
	}})
}
