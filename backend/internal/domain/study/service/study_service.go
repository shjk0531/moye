package service

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/chat/channel/model"
	studyModel "github.com/shjk0531/moye/backend/internal/domain/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
	"gorm.io/gorm"
)

type StudyService interface {
	CreateStudy(study *studyModel.Study, userID uuid.UUID) error
	GetStudy(id uuid.UUID) (*studyModel.Study, error)
	GetAllStudies() ([]*studyModel.Study, error)
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

func (s *studyService) CreateStudy(study *studyModel.Study, userID uuid.UUID) error {
	// 트랜잭션 시작
	err := s.db.Transaction(func(tx *gorm.DB) error {
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