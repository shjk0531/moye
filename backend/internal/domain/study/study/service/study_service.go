package service

import (
	"context"

	"github.com/google/uuid"
	channelRepository "github.com/shjk0531/moye/backend/internal/domain/study/channel/repository"
	channelService "github.com/shjk0531/moye/backend/internal/domain/study/channel/service"

	"github.com/shjk0531/moye/backend/internal/domain/study/study/dto"
	studyModel "github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/repository"
	"gorm.io/gorm"
)

type StudyService interface {
	CreateStudy(ctx context.Context, study *studyModel.Study, userID uuid.UUID) error
	GetStudy(id uuid.UUID) (*studyModel.Study, error)
	GetAllStudies() ([]*studyModel.Study, error)
	GetUserStudies(userID uuid.UUID) ([]*studyModel.Study, error)
	GetSimpleStudyList(page, size int) (dto.SimpleStudyListResponse, error)
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

        orders, err := channelSvc.GetChannelOrders(studyID)
        if err != nil {
            return err
        }

        nextPosition := 0
        if len(orders) > 0 {
            nextPosition = orders[len(orders)-1].Position + 1
        }
        
        generalCh, err := channelSvc.CreateChannel(studyID, "일반", nextPosition +2)
        if err != nil {
            return err
        }
        announcementCh, err := channelSvc.CreateChannel(studyID, "공지사항", nextPosition)
        if err != nil {
            return err
        }

        // 6) 기본 채널 그룹 생성
        defaultGroup, err := channelSvc.CreateChannelGroup(studyID, "기본", nextPosition + 1)
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


// GetSimpleStudyList godoc
// @Summary 스터디 목록 조회
// @Description 스터디 목록을 조회
// @Tags studies
// @Accept json
// @Produce json
// @Param page query int false "페이지 번호"
// @Param size query int false "페이지 크기"
func (s *studyService) GetSimpleStudyList(page, size int) (dto.SimpleStudyListResponse, error) {
    offset := (page - 1) * size 
    simpleStudies, err := s.repo.GetSimpleStudyList(offset, size)
    if err != nil {
        return dto.SimpleStudyListResponse{}, err
    }

    return dto.SimpleStudyListResponse{
        Studies: simpleStudies,
    }, nil
}
