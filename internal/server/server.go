package server

import (
	"fmt"
	"hello-go/internal/data"
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
	data := data.NewData(mysql, nil)
	service := service.NewService(data)
	auth := s.Engine.Group("/auth")
	{
		//curl -v -X POST -d "username=admin&password=123456" http://localhost:8080/auth/login
		auth.POST("/login", service.Login)
		// curl -v -X POST -d "username=admin&password=123456" http://localhost:8080/auth/register
		auth.POST("/register", service.Register)
	}
	return s.Engine.Run(":8080")
}

func NewServer() *Server {
	return &Server{gin.New()}
}
