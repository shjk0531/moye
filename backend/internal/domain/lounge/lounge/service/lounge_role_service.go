package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/repository"
	globalModel "github.com/shjk0531/moye/backend/internal/global/model"
)

type LoungeRoleService interface {
	CreateLoungeRole(loungeMemberID uuid.UUID, name string, colorHex string) (uuid.UUID, error)
}

type loungeRoleService struct {
	repo repository.Repository
}

func NewLoungeRoleService(repo repository.Repository) LoungeRoleService {
	return &loungeRoleService{repo: repo}
}

// 역할 생성
func (s *loungeRoleService) CreateLoungeRole(loungeMemberID uuid.UUID, role_name string, colorHex string) (uuid.UUID, error) {
	role := &model.LoungeMemberRole{
		LoungeMemberID: loungeMemberID,
		RoleName:          role_name,
		ColorHex:      colorHex,
	}
	return s.repo.CreateRole(role)
}

// 역할 업데이트
func (s *loungeRoleService) UpdateLoungeRole(loungeRoleID uuid.UUID, role_name string, colorHex string) error {
	role := &model.LoungeMemberRole{
		BaseModel: globalModel.BaseModel{
			ID: loungeRoleID,
		},
		RoleName:     role_name,
		ColorHex: colorHex,
	}
	return s.repo.UpdateRole(role)
}

// 역할 삭제
