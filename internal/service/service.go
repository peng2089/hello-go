package service

type Service struct {
	As *AuthService
}

func NewService(as *AuthService) *Service {
	return &Service{As: as}
}
