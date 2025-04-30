package service

import (
	"context"

	"github.com/google/uuid"
	channelRepository "github.com/shjk0531/moye/backend/internal/domain/chat/channel/repository"
	channelService "github.com/shjk0531/moye/backend/internal/domain/chat/channel/service"
	"github.com/shjk0531/moye/backend/internal/domain/study/dto"
	studyModel "github.com/shjk0531/moye/backend/internal/domain/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
	"gorm.io/gorm"
)

type StudyService interface {
	CreateStudy(ctx context.Context, study *studyModel.Study, userID uuid.UUID) error
	GetStudy(id uuid.UUID) (*studyModel.Study, error)
	GetAllStudies() ([]*studyModel.Study, error)
	GetUserStudies(userID uuid.UUID) ([]*studyModel.Study, error)
	GetUserStudyWithRole(studyID, userID uuid.UUID) (*dto.UserStudyDTO, error)
	GetUserStudiesWithRoles(userID uuid.UUID) (*dto.UserStudiesResponse, error)
}

type studyService struct {
	repo repository.Repository
	db   *gorm.DB
}

func NewStudyService(repo repository.Repository, db *gorm.DB) StudyService {
	return &studyService{
		repo: repo,
		db:   db,
	}
}

func (s *studyService) CreateStudy(ctx context.Context, study *studyModel.Study, userID uuid.UUID) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // 1) Transaction 전용 StudyRepository
        studyRepo := repository.NewRepository(tx)

        // 2) Transaction 전용 ChannelRepository / Service
        chRepo := channelRepository.NewChannelRepository(tx)
        chSvc := channelService.NewChannelService(chRepo)

        // —————————————————————————————————————————————
        // 3) 스터디 생성
        studyID, err := studyRepo.CreateStudy(study)
        if err != nil {
            return err
        }

        // 4) “관리자” 롤 생성
        adminRoleID, err := studyRepo.CreateRole(&studyModel.StudyRole{
            Name:     "관리자",
            ColorHex: "#000000",
        })
        if err != nil {
            return err
        }

        // 5) 스터디 멤버(관리자) 등록
        if err := studyRepo.CreateStudyMember(&studyModel.StudyMember{
            StudyID: studyID,
            UserID:  userID,
            RoleID:  adminRoleID,
        }); err != nil {
            return err
        }

        // —————————————————————————————————————————————
        // 6) 기본 채널 생성 (“일반”, “공지사항”)
        generalCh, err := chSvc.CreateChannel(studyID, "일반", 0)
        if err != nil {
            return err
        }
        announcementCh, err := chSvc.CreateChannel(studyID, "공지사항", 1)
        if err != nil {
            return err
        }

        // 7) 기본 채널 그룹 생성 (position 0)
        defaultGroup, err := chSvc.CreateChannelGroup(studyID, "기본", 0)
        if err != nil {
            return err
        }

        // 8) 그룹에 채널 순서대로 추가
        if err := chSvc.AddChannelToGroup(ctx, defaultGroup.ID, generalCh.ID, 0); err != nil {
            return err
        }
        if err := chSvc.AddChannelToGroup(ctx, defaultGroup.ID, announcementCh.ID, 1); err != nil {
            return err
        }

        // 트랜잭션 커밋
        return nil
    })
}

// 역할 생성
func (s *studyService) CreateRole(name string, colorHex string) (uuid.UUID, error) {
	return s.repo.CreateRole(&studyModel.StudyRole{
		Name:     name,
		ColorHex: colorHex,
	})
}

// 스터디 멤버 생성
func (s *studyService) CreateStudyMember(studyID uuid.UUID, userID uuid.UUID, nickname string, profileURL string, roleID uuid.UUID) error {
	return s.repo.CreateStudyMember(&studyModel.StudyMember{
		StudyID:    studyID,
		UserID:     userID,
		Nickname:   nickname,
		ProfileURL: profileURL,
		RoleID:     roleID,
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

// GetUserStudyWithRole은 특정 스터디와 사용자의 역할 정보를 함께 반환합니다.
func (s *studyService) GetUserStudyWithRole(studyID, userID uuid.UUID) (*dto.UserStudyDTO, error) {
	// 1. 스터디 정보 조회
	study, err := s.repo.FindByID(studyID)
	if err != nil {
		return nil, err
	}

	// 2. 사용자의 스터디 멤버 정보 조회
	var member studyModel.StudyMember
	if err := s.db.Where("study_id = ? AND user_id = ?", studyID, userID).First(&member).Error; err != nil {
		// 멤버가 존재하지 않는 경우는 역할 정보 없이 반환
		studyDTO := &dto.UserStudyDTO{
			ID:          study.ID,
			Name:        study.Name,
			ProfileURL:  study.ProfileURL,
			Description: study.Description,
			Tags:        study.Tags,
			CreatedAt:   study.CreatedAt,
			UpdatedAt:   study.UpdatedAt,
		}
		return studyDTO, nil
	}

	// 3. 역할 정보가 있는 경우 역할 조회
	var role *studyModel.StudyRole
	if member.RoleID != uuid.Nil {
		role = new(studyModel.StudyRole)
		if err := s.db.Where("id = ?", member.RoleID).First(role).Error; err != nil {
			// 역할이 존재하지 않는 경우는 무시
			role = nil
		}
	}

	// 4. DTO 변환
	studyDTO := &dto.UserStudyDTO{
		ID:          study.ID,
		Name:        study.Name,
		ProfileURL:  study.ProfileURL,
		Description: study.Description,
		Tags:        study.Tags,
		CreatedAt:   study.CreatedAt,
		UpdatedAt:   study.UpdatedAt,
	}

	// 역할 정보가 있으면 추가
	if role != nil {
		studyDTO.UserRole = &dto.UserRoleDTO{
			ID:        role.ID,
			Name:      role.Name,
			ColorHex:  role.ColorHex,
			RoleFlags: member.RoleFlags,
		}
	}

	return studyDTO, nil
}

// GetUserStudiesWithRoles는 사용자가 속한 모든 스터디의 목록과 각 스터디에서의 역할 정보를 반환합니다.
func (s *studyService) GetUserStudiesWithRoles(userID uuid.UUID) (*dto.UserStudiesResponse, error) {
	// 1. 사용자가 속한 스터디 목록 조회
	studies, err := s.repo.FindStudiesByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 2. 응답 DTO 초기화
	response := &dto.UserStudiesResponse{
		Studies: make([]dto.StudyListDTO, 0, len(studies)),
	}

	// 3. 각 스터디별로 study_member의 updated_at 조회 및 DTO 구성
	for _, study := range studies {
		// 해당 사용자의 study_member 레코드 조회
		var member studyModel.StudyMember
		if err := s.db.Where("study_id = ? AND user_id = ?", study.ID, userID).First(&member).Error; err != nil {
			// 조회 실패 시 스터디의 updatedAt을 대체값으로 사용
			member.UpdatedAt = study.UpdatedAt
		}
		response.Studies = append(response.Studies, dto.StudyListDTO{
			ID:          study.ID,
			Name:        study.Name,
			ProfileURL:  study.ProfileURL,
			UpdatedAt:   member.UpdatedAt,
		})
	}

	return response, nil
}