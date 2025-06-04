package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/repository"
)

type LoungeMemberService interface {
	AddLoungeMember(loungeID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error
}

type loungeMemberService struct {
	repo repository.Repository
}

func NewLoungeMemberService(repo repository.Repository) LoungeMemberService {
	return &loungeMemberService{repo: repo}
}

// 라운지에 멤버 추가
func (s *loungeMemberService) AddLoungeMember(loungeID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	member := &model.LoungeMember{
		LoungeID: loungeID,
		UserID:  userID,
		RoleID:  roleID,
	}
	return s.repo.CreateLoungeMember(member)
}
