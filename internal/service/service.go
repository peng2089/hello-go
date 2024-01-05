package service

import "hello-go/internal/data"

type Service struct {
	Data *data.Data
}

func NewService(data *data.Data) *Service {
	return &Service{Data: data}
}
