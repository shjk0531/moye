package service

import (
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
	"gorm.io/gorm"
)

type Service interface {
	GetStudyService() StudyService
	GetStudyMemberService() StudyMemberService
	GetStudyRoleService() StudyRoleService
}

type service struct {
	repo repository.Repository
	db *gorm.DB
	studyService StudyService
	studyMemberService StudyMemberService
	studyRoleService StudyRoleService
}

func NewService(repo repository.Repository, db *gorm.DB) Service {
	return &service{
		repo: repo,
		db: db,
	}
}

// Init 함수는 모든 서비스를 초기화합니다
func Init(repo repository.Repository, db *gorm.DB) Service {
	svc := &service{
		repo: repo,
		db: db,
	}
	
	// 각 서비스 초기화
	svc.studyRoleService = NewStudyRoleService(repo)
	svc.studyMemberService = NewStudyMemberService(repo)
	svc.studyService = NewStudyService(repo, db, svc.studyRoleService, svc.studyMemberService)
	
	return svc
}

func (s *service) GetStudyService() StudyService {
	return s.studyService
}

func (s *service) GetStudyMemberService() StudyMemberService {
	return s.studyMemberService
}

func (s *service) GetStudyRoleService() StudyRoleService {
	return s.studyRoleService
}