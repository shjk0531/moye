package repository

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(study *model.Study) error
	FindByID(id uuid.UUID) (*model.Study, error)
	FindAll() ([]*model.Study, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(study *model.Study) error {
	return r.db.Create(study).Error
}

func (r *repository) FindByID(id uuid.UUID) (*model.Study, error) {
	var study model.Study
	if err := r.db.First(&study, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &study, nil
}

func (r *repository) FindAll() ([]*model.Study, error) {
	var studies []*model.Study
	if err := r.db.Find(&studies).Error; err != nil {
		return nil, err
	}
	return studies, nil
}