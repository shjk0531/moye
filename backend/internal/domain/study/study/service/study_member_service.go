package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/repository"
)

type StudyMemberService interface {
	AddStudyMember(studyID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error
}

type studyMemberService struct {
	repo repository.Repository
}

func NewStudyMemberService(repo repository.Repository) StudyMemberService {
	return &studyMemberService{repo: repo}
}

// 스터디에 멤버 추가
func (s *studyMemberService) AddStudyMember(studyID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	member := &model.StudyMember{
		StudyID: studyID,
		UserID:  userID,
		RoleID:  roleID,
	}
	return s.repo.CreateStudyMember(member)
}
