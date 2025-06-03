package service

import (
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/repository"
	"gorm.io/gorm"
)

type Service interface {
	GetLoungeService() LoungeService
	GetLoungeMemberService() LoungeMemberService
	GetLoungeRoleService() LoungeRoleService
}

type service struct {
	repo repository.Repository
	db *gorm.DB
	loungeService LoungeService
	loungeMemberService LoungeMemberService
	loungeRoleService LoungeRoleService
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
	svc.loungeRoleService = NewLoungeRoleService(repo)
	svc.loungeMemberService = NewLoungeMemberService(repo)
	svc.loungeService = NewLoungeService(repo, db, svc.loungeRoleService, svc.loungeMemberService)
	
	return svc
}

func (s *service) GetLoungeService() LoungeService {
	return s.loungeService
}

func (s *service) GetLoungeMemberService() LoungeMemberService {
	return s.loungeMemberService
}

func (s *service) GetLoungeRoleService() LoungeRoleService {
	return s.loungeRoleService
}