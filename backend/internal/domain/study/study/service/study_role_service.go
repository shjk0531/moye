package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/repository"
	globalModel "github.com/shjk0531/moye/backend/internal/global/model"
)

type StudyRoleService interface {
	CreateStudyRole(studyMemberID uuid.UUID, name string, colorHex string) (uuid.UUID, error)
}

type studyRoleService struct {
	repo repository.Repository
}

func NewStudyRoleService(repo repository.Repository) StudyRoleService {
	return &studyRoleService{repo: repo}
}

// 역할 생성
func (s *studyRoleService) CreateStudyRole(studyMemberID uuid.UUID, name string, colorHex string) (uuid.UUID, error) {
	role := &model.StudyMemberRole{
		StudyMemberID: studyMemberID,
		Name:          name,
		ColorHex:      colorHex,
	}
	return s.repo.CreateRole(role)
}

// 역할 업데이트
func (s *studyRoleService) UpdateStudyRole(studyRoleID uuid.UUID, name string, colorHex string) error {
	role := &model.StudyMemberRole{
		BaseModel: globalModel.BaseModel{
			ID: studyRoleID,
		},
		Name:     name,
		ColorHex: colorHex,
	}
	return s.repo.UpdateRole(role)
}

// 역할 삭제
