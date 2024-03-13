package service

import (
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

func (s *UserService) Info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
