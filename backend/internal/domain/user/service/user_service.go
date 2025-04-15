package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
)

type Service interface {
	RegisterUser(user *model.User) error
	GetUser(id uuid.UUID) (*model.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *service) GetUser(id uuid.UUID) (*model.User, error) {
	return s.repo.FindByID(id)
}
