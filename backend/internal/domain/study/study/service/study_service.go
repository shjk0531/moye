package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/repository"
)

type Service interface {
	CreateStudy(study *model.Study) error
	GetStudy(id uuid.UUID) (*model.Study, error)
	GetAllStudies() ([]*model.Study, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStudy(study *model.Study) error {
	return s.repo.Create(study)
}

func (s *service) GetStudy(id uuid.UUID) (*model.Study, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetAllStudies() ([]*model.Study, error) {
	return s.repo.FindAll()
}