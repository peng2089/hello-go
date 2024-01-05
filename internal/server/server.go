package server

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
}

func (s *Server) Run() error {
	return s.Engine.Run(":8080")
}

func NewServer() *Server {
	return &Server{gin.New()}
}
