package repository

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/global/constants"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByID(id uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindUserWithRolesAndPermissions(id uuid.UUID) (*model.User, error)
	GetUserRoles(userID uuid.UUID) ([]constants.UserRoleName, error)
	GetUserPermissions(userID uuid.UUID) ([]constants.UserPermissionName, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserWithRolesAndPermissions는 사용자와 관련된 역할 및 권한 정보를 포함하여 조회
func (r *repository) FindUserWithRolesAndPermissions(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserRoles는 사용자의 역할 목록을 조회
func (r *repository) GetUserRoles(userID uuid.UUID) ([]constants.UserRoleName, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	// pq.StringArray에서 []constants.UserRoleName으로 변환
	roles := make([]constants.UserRoleName, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = constants.UserRoleName(r)
	}
	
	return roles, nil
}

// GetUserPermissions는 사용자의 권한 목록을 조회
func (r *repository) GetUserPermissions(userID uuid.UUID) ([]constants.UserPermissionName, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	
	// pq.StringArray에서 []constants.UserPermissionName으로 변환
	permissions := make([]constants.UserPermissionName, len(user.Permissions))
	for i, p := range user.Permissions {
		permissions[i] = constants.UserPermissionName(p)
	}
	
	return permissions, nil
}
