package repository

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id uuid.UUID) (*model.Study, error)
	FindAll() ([]*model.Study, error)
	FindStudiesByUserID(userID uuid.UUID) ([]*model.Study, error)
	CreateRole(role *model.StudyRole) (uuid.UUID, error)
	FindRoleByID(id uuid.UUID) (*model.StudyRole, error)
	UpdateRole(role *model.StudyRole) error
	CreateStudyMember(member *model.StudyMember) error
	CreateStudy(study *model.Study) (uuid.UUID, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
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

// 역할 생성
func (r *repository) CreateRole(role *model.StudyRole) (uuid.UUID, error) {
	err := r.db.Create(role).Error
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// 역할 조회
func (r *repository) FindRoleByID(id uuid.UUID) (*model.StudyRole, error) {
	var role model.StudyRole
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// 역할 업데이트
func (r *repository) UpdateRole(role *model.StudyRole) error {
	return r.db.Save(role).Error
}

// 스터디 멤버 생성
func (r *repository) CreateStudyMember(member *model.StudyMember) error {
	return r.db.Create(member).Error
}

// 스터디 생성
func (r *repository) CreateStudy(study *model.Study) (uuid.UUID, error) {
	err := r.db.Create(study).Error
	if err != nil {
		return uuid.Nil, err
	}
	return study.ID, nil
}

// 스터디 조회

