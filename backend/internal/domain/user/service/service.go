package service

import "github.com/shjk0531/moye/backend/internal/domain/user/repository"

type Service struct {
	User UserService
	Auth AuthService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
		Auth: NewAuthService(repo),
	}
}