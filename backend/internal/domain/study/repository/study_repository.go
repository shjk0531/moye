package repository

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(study *model.Study) error
	FindByID(id uuid.UUID) (*model.Study, error)
	FindAll() ([]*model.Study, error)
	FindStudiesByUserID(userID uuid.UUID) ([]*model.Study, error)
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

// FindStudiesByUserID는 특정 사용자가 속한 모든 스터디 목록을 조회합니다.
func (r *repository) FindStudiesByUserID(userID uuid.UUID) ([]*model.Study, error) {
	var studies []*model.Study
	
	// study_members 테이블을 통해 사용자가 속한 스터디 ID를 조회하고,
	// 해당 ID로 스터디 정보를 가져옵니다.
	err := r.db.Table("studies").
		Joins("INNER JOIN study_members ON studies.id = study_members.study_id").
		Where("study_members.user_id = ?", userID).
		Find(&studies).Error
	
	if err != nil {
		return nil, err
	}
	
	return studies, nil
}