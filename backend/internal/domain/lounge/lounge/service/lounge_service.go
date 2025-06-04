package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	channelRepository "github.com/shjk0531/moye/backend/internal/domain/lounge/channel/repository"
	channelService "github.com/shjk0531/moye/backend/internal/domain/lounge/channel/service"

	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/dto"
	loungeModel "github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/repository"
	"gorm.io/gorm"
)

type LoungeService interface {
	CreateLounge(ctx context.Context, lounge *loungeModel.Lounge, userID uuid.UUID) error
	GetLounge(id uuid.UUID) (*loungeModel.Lounge, error)
	GetLoungeList() (dto.LoungeListResponse, error)
	GetUserLounges(userID uuid.UUID) (dto.LoungeIconResponse, error)
}

type loungeService struct {
	repo       repository.Repository
	db         *gorm.DB
	roleService LoungeRoleService
	memberService LoungeMemberService
}

func NewLoungeService(repo repository.Repository, db *gorm.DB, roleService LoungeRoleService, memberService LoungeMemberService) LoungeService {
	return &loungeService{
		repo:          repo,
		db:            db,
		roleService:   roleService,
		memberService: memberService,
	}
}

func (s *loungeService) CreateLounge(ctx context.Context, lounge *loungeModel.Lounge, userID uuid.UUID) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // 1) Transaction 전용 ChannelRepository / Service
        channelRepo := channelRepository.NewChannelRepository(tx)
        channelSvc := channelService.NewChannelService(channelRepo)

        // —————————————————————————————————————————————
        // 2) 라운지 생성
        loungeID, err := s.repo.CreateLounge(lounge)
        if err != nil {
            fmt.Println("CreateLounge error", err)
            return err
        }

		// 3) "관리자" 롤 생성
		roleID, err := s.roleService.CreateLoungeRole(userID, "관리자", "#000000")
		if err != nil {
            fmt.Println("CreateLoungeRole error", err)
			return err
		}


        // 4) 라운지 멤버(관리자) 등록
        if err := s.memberService.AddLoungeMember(loungeID, userID, roleID); err != nil {
            fmt.Println("AddLoungeMember error", err)
            return err
        }

        // 그룹 없이 채널 생성
        _, _, err = channelSvc.CreateChannel(ctx, loungeID, "일반 채널", nil)
        if err != nil {
            fmt.Println("CreateChannel error", err)
            return err
        }
        _, _, err = channelSvc.CreateChannel(ctx, loungeID, "공지사항 채널", nil)
        if err != nil {
            fmt.Println("CreateChannel error", err)
            return err
        }


        // 6) 기본 채널 그룹 생성
        group, err := channelSvc.CreateChannelGroup(ctx, loungeID, "기본")
        if err != nil {
            fmt.Println("CreateChannelGroup error", err)
            return err
        }


        _, _, err = channelSvc.CreateChannel(ctx, loungeID, "그룹 채널", &group.ID)
        if err != nil {
            fmt.Println("CreateChannel error", err)
            return err
        }

        return nil
    })
}

func (s *loungeService) GetLounge(id uuid.UUID) (*loungeModel.Lounge, error) {
	return s.repo.FindByID(id)
}

func (s *loungeService) GetLoungeList() (dto.LoungeListResponse, error) {
	lounges, err := s.repo.FindLoungeList()
	if err != nil {
		return dto.LoungeListResponse{}, err
	}
	return dto.LoungeListResponse{
		Lounges: lounges,
	}, nil
}

// 특정 사용자가 속한 모든 라운지 목록을 반환
func (s *loungeService) GetUserLounges(userID uuid.UUID) (dto.LoungeIconResponse, error) {
	icons, err := s.repo.FindLoungesByUserID(userID)
	if err != nil {
		return dto.LoungeIconResponse{}, err
	}
	return dto.LoungeIconResponse{
		Icons: icons,
	}, nil
}
