package service

import (
	"context"

	"github.com/google/uuid"
	channelRepository "github.com/shjk0531/moye/backend/internal/domain/study/channel/repository"
	channelService "github.com/shjk0531/moye/backend/internal/domain/study/channel/service"

	studyModel "github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/repository"
	"gorm.io/gorm"
)

type StudyService interface {
	CreateStudy(ctx context.Context, study *studyModel.Study, userID uuid.UUID) error
	GetStudy(id uuid.UUID) (*studyModel.Study, error)
	GetAllStudies() ([]*studyModel.Study, error)
	GetUserStudies(userID uuid.UUID) ([]*studyModel.Study, error)
}

type studyService struct {
	repo       repository.Repository
	db         *gorm.DB
	roleService StudyRoleService
	memberService StudyMemberService
}

func NewStudyService(repo repository.Repository, db *gorm.DB, roleService StudyRoleService, memberService StudyMemberService) StudyService {
	return &studyService{
		repo:          repo,
		db:            db,
		roleService:   roleService,
		memberService: memberService,
	}
}

func (s *studyService) CreateStudy(ctx context.Context, study *studyModel.Study, userID uuid.UUID) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // 1) Transaction 전용 ChannelRepository / Service
        channelRepo := channelRepository.NewChannelRepository(tx)
        channelSvc := channelService.NewChannelService(channelRepo)

        // —————————————————————————————————————————————
        // 2) 스터디 생성
        studyID, err := s.repo.CreateStudy(study)
        if err != nil {
            return err
        }

		// 3) "관리자" 롤 생성
		roleID, err := s.roleService.CreateStudyRole(userID, "관리자", "#000000")
		if err != nil {
			return err
		}


        // 4) 스터디 멤버(관리자) 등록
        if err := s.memberService.AddStudyMember(studyID, userID, roleID); err != nil {
            return err
        }

        // —————————————————————————————————————————————
        // 5) 기본 채널 생성 ("일반", "공지사항")
        generalCh, err := channelSvc.CreateChannel(studyID, "일반", 0)
        if err != nil {
            return err
        }
        announcementCh, err := channelSvc.CreateChannel(studyID, "공지사항", 1)
        if err != nil {
            return err
        }

        // 6) 기본 채널 그룹 생성 (position 0)
        defaultGroup, err := channelSvc.CreateChannelGroup(studyID, "기본", 0)
        if err != nil {
            return err
        }

        // 7) 그룹에 채널 순서대로 추가
        if err := channelSvc.AddChannelToGroup(ctx, defaultGroup.ID, generalCh.ID, 0); err != nil {
            return err
        }
        if err := channelSvc.AddChannelToGroup(ctx, defaultGroup.ID, announcementCh.ID, 1); err != nil {
            return err
        }

        return nil
    })
}

func (s *studyService) GetStudy(id uuid.UUID) (*studyModel.Study, error) {
	return s.repo.FindByID(id)
}

func (s *studyService) GetAllStudies() ([]*studyModel.Study, error) {
	return s.repo.FindAll()
}

// GetUserStudies는 특정 사용자가 속한 모든 스터디 목록을 반환합니다.
func (s *studyService) GetUserStudies(userID uuid.UUID) ([]*studyModel.Study, error) {
	return s.repo.FindStudiesByUserID(userID)
}
