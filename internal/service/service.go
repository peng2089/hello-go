package service

type Service struct {
	As *AuthService
	Us *UserService
}

func NewService(as *AuthService, us *UserService) *Service {
	return &Service{As: as, Us: us}
}
