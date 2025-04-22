package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
)

type UserService interface {
	RegisterUser(user *model.User) error
	GetUser(id uuid.UUID) (*model.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUser(id uuid.UUID) (*model.User, error) {
	return s.repo.FindByID(id)
}
