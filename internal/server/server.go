package server

import (
	"fmt"
	"hello-go/internal/biz"
	"hello-go/internal/data"
	"hello-go/internal/pkg/middleware"
	"hello-go/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func (s *Server) Run() error {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	mysql, err := data.NewMysql()
	if err != nil {
		fmt.Printf("Err: %+v\n", err)
		return err
	}
	d := data.NewData(mysql, nil)
	repo := data.NewUserRepo(d)
	userUsecase := biz.NewUserUsecase(repo)
	authService := service.NewAuthService(userUsecase)
	userService := service.NewUserService(userUsecase)
	service := service.NewService(authService, userService)

	auth := s.Engine.Group("/auth")
	{
		// curl -v -d "username=admin&password=123456" -X POST http://localhost:8080/auth/login
		auth.POST("/login", service.As.Login)
		// curl -v -X POST -d "username=admin&password=123456" http://localhost:8080/auth/register
		auth.POST("/register", service.As.Register)
	}
	user := s.Engine.Group("/user", middleware.TokenMiddleware(userUsecase))
	{
		user.GET("/info", service.Us.Info)
	}
	return s.Engine.Run(":8080")
}

func NewServer() *Server {
	return &Server{gin.New()}
}
