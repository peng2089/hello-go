package service

import (
	"fmt"
	"hello-go/internal/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// Info 获取用户信息
func (s *UserService) Info(ctx *gin.Context) {
	userCtx, exists := ctx.Get("user")
	if !exists {
		fmt.Println("user 不存在")
	}
	user := userCtx.(*biz.User)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    user,
	})
}
