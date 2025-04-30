package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/chat/channel/model"
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
	// 트랜잭션 시작 (Gin context를 전달)
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 스터디 생성
		if err := tx.Create(study).Error; err != nil {
			return err
		}

		// 2. 기본 역할 생성 - 리더 역할
		leaderRole := &studyModel.StudyRole{
			Name:     "리더",
			ColorHex: "#FF5733", // 기본 리더 색상
		}

		// 3. 스터디 멤버 생성
		member := &studyModel.StudyMember{
			StudyID:    study.ID,
			UserID:     userID,
			Nickname:   "리더", // 기본 닉네임, 나중에 사용자 이름으로 변경 가능
			ProfileURL: "",    // 기본 프로필 URL
			RoleFlags:  1,     // 리더 권한 부여
		}

		if err := tx.Create(member).Error; err != nil {
			return err
		}

		// 생성된 멤버 ID를 역할에 설정
		leaderRole.StudyMemberID = member.ID
		if err := tx.Create(leaderRole).Error; err != nil {
			return err
		}

		// 멤버의 역할 ID 업데이트
		member.RoleID = leaderRole.ID
		if err := tx.Save(member).Error; err != nil {
			return err
		}

		// 4. 기본 채널 그룹 생성
		channelGroup := &model.ChannelGroup{
			StudyID: study.ID,
			Name:    "일반",
		}
		if err := tx.Create(channelGroup).Error; err != nil {
			return err
		}

		// 5. 기본 채널 생성 - 일반 채널
		generalChannel := &model.Channel{
			StudyID: study.ID,
			Name:    "일반",
		}
		if err := tx.Create(generalChannel).Error; err != nil {
			return err
		}

		// 공지사항 채널
		announcementChannel := &model.Channel{
			StudyID: study.ID,
			Name:    "공지사항",
		}
		if err := tx.Create(announcementChannel).Error; err != nil {
			return err
		}

		// 6. 채널 그룹 내 채널 순서 설정
		generalChannelOrder := &model.ChannelGroupOrder{
			GroupID:   channelGroup.ID,
			Position:  0,
			ChannelID: generalChannel.ID,
		}
		if err := tx.Create(generalChannelOrder).Error; err != nil {
			return err
		}

		announcementChannelOrder := &model.ChannelGroupOrder{
			GroupID:   channelGroup.ID,
			Position:  1,
			ChannelID: announcementChannel.ID,
		}
		if err := tx.Create(announcementChannelOrder).Error; err != nil {
			return err
		}

		// 7. 채널 그룹 순서 설정
		channelGroupOrder := &model.ChannelOrder{
			StudyID:  study.ID,
			Position: 0,
			ItemType: "group",
			ItemID:   channelGroup.ID,
		}
		if err := tx.Create(channelGroupOrder).Error; err != nil {
			return err
		}

		return nil
	})

	return err
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
		Studies: make([]dto.UserStudyDTO, 0, len(studies)),
	}

	// 3. 각 스터디별로 역할 정보 조회 및 DTO 구성
	for _, study := range studies {
		studyDTO, err := s.GetUserStudyWithRole(study.ID, userID)
		if err != nil {
			// 오류 발생 시 해당 스터디는 건너뛰기
			continue
		}
		response.Studies = append(response.Studies, *studyDTO)
	}

	return response, nil
}