package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hello-go/internal/data"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20" example:"test"`
	Password string `form:"password" json:"password" binding:"required,min=6" example:"123456"`
}

type LoginResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20" example:"test"`
	Password string `form:"password" json:"password" binding:"required,min=6" example:"123456"`
}
type RegisterResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

// Login: 登录
func (s *Service) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("req: %+v\n", req)

	var user data.UserLocalAuth
	// 1. 查找用户
	ret := s.Data.Db.Where("username = ?", req.Username).First(&user)
	if ret.Error != nil {
		if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "用户不存在",
			})
			return
		}
	}
	// 2. 比对密码
	oldPassword := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	if strings.Compare(user.Password, oldPassword) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "密码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &LoginResp{
		ID:       user.UserId,
		Username: user.Username,
	}})
}

// Register: 注册
func (s *Service) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user data.User
	var userLocalAuth data.UserLocalAuth

	// 1. 检测用户名是否已存在
	ret := s.Data.Db.Where("username = ?", req.Username).First(&userLocalAuth)
	if ret.Error != nil {
		if !errors.Is(ret.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": ret.Error,
			})
			return
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "用户已存在",
		})
		return
	}
	// 2. 添加用户
	s.Data.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err // roll back
		}
		userLocalAuth = data.UserLocalAuth{
			UserId:   user.ID,
			Username: req.Username,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
			// CreatedAt: time.Now().Unix(),
		}
		if err := tx.Create(&userLocalAuth).Error; err != nil {
			return err // roll back
		}

		// 返回nil，提交事务
		return nil
	})
	ctx.JSON(http.StatusOK, gin.H{"data": &RegisterResp{
		ID:       user.ID,
		Username: userLocalAuth.Username,
	}})
}
