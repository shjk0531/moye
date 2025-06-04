package repository

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/dto"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id uuid.UUID) (*model.Lounge, error)
	FindLoungeList() ([]dto.LoungeListDTO, error)
	FindLoungesByUserID(userID uuid.UUID) ([]dto.LoungeIconDTO, error)
	CreateRole(role *model.LoungeMemberRole) (uuid.UUID, error)
	FindRoleByID(id uuid.UUID) (*model.LoungeMemberRole, error)
	UpdateRole(role *model.LoungeMemberRole) error
	CreateLoungeMember(member *model.LoungeMember) error
	CreateLounge(lounge *model.Lounge) (uuid.UUID, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}


func (r *repository) FindByID(id uuid.UUID) (*model.Lounge, error) {
	var lounge model.Lounge
	if err := r.db.First(&lounge, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &lounge, nil
}

func (r *repository) FindLoungeList() ([]dto.LoungeListDTO, error) {
	// 라운지 목록 조회
	var dtos []dto.LoungeListDTO
	if err := r.db.
		Model(&model.Lounge{}).
		Scan(&dtos).Error; err != nil {
		return nil, err
	}
	return dtos, nil
}

// FindLoungesByUserID는 특정 사용자가 속한 모든 라운지 목록을 조회합니다.
func (r *repository) FindLoungesByUserID(userID uuid.UUID) ([]dto.LoungeIconDTO, error) {
	var dtos []dto.LoungeIconDTO
	
	// lounge_members 테이블을 통해 사용자가 속한 라운지 ID를 조회하고,
	// 해당 ID로 라운지 정보를 가져옵니다.
	err := r.db.Table("lounges").
		Joins("INNER JOIN lounge_members ON lounges.id = lounge_members.lounge_id").
		Where("lounge_members.user_id = ?", userID).
		Select("lounges.id", "lounges.name", "lounges.profile_url").
		Find(&dtos).Error
	
	if err != nil {
		return nil, err
	}
	
	return dtos, nil
}

// 역할 생성
func (r *repository) CreateRole(role *model.LoungeMemberRole) (uuid.UUID, error) {
	err := r.db.Create(role).Error
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// 역할 조회
func (r *repository) FindRoleByID(id uuid.UUID) (*model.LoungeMemberRole, error) {
	var role model.LoungeMemberRole
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// 역할 업데이트
func (r *repository) UpdateRole(role *model.LoungeMemberRole) error {
	return r.db.Save(role).Error
}

// 라운지 멤버 생성
func (r *repository) CreateLoungeMember(member *model.LoungeMember) error {
	return r.db.Create(member).Error
}

// 라운지 생성
func (r *repository) CreateLounge(lounge *model.Lounge) (uuid.UUID, error) {
	err := r.db.Create(lounge).Error
	if err != nil {
		return uuid.Nil, err
	}
	return lounge.ID, nil
}
