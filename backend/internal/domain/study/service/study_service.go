package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
)

type StudyService interface {
	CreateStudy(study *model.Study) error
	GetStudy(id uuid.UUID) (*model.Study, error)
	GetAllStudies() ([]*model.Study, error)
}

type sutdyService struct {
	repo repository.Repository
}

func NewStudyService(repo repository.Repository) StudyService {
	return &sutdyService{repo: repo}
}

func (s *sutdyService) CreateStudy(study *model.Study) error {
	return s.repo.Create(study)
}

func (s *sutdyService) GetStudy(id uuid.UUID) (*model.Study, error) {
	return s.repo.FindByID(id)
}

func (s *sutdyService) GetAllStudies() ([]*model.Study, error) {
	return s.repo.FindAll()
}