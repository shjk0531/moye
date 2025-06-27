package dto

import (
	"time"

	"github.com/google/uuid"
)

type LoungeIconDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProfileURL  string    `json:"profile_url"`
}

// lounge 목록 조회 시 사용되는 DTO
type LoungeListDTO struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	ProfileURL  string    `json:"profile_url" gorm:"column:profile_url"`
	Description string    `json:"description" gorm:"column:description"`
	Tags        []string  `json:"tags" gorm:"column:tags;type:jsonb;serializer:json"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}


// UserLoungeDTO는 사용자의 라운지 정보를 표현하는 DTO
type UserLoungeDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProfileURL  string    `json:"profile_url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	// 사용자 역할 정보
	UserRole    *UserRoleDTO `json:"user_role,omitempty"`
}


// UserRoleDTO는 해당 라운지에서 사용자의 역할 정보를 표현하는 DTO입니다.
type UserRoleDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ColorHex  string    `json:"color_hex"`
	RoleFlags int64     `json:"role_flags"`
}

// LoungeMemberInfoDTO 라운지 멤버의 상세 정보를 담는 DTO
type LoungeMemberInfoDTO struct {
    UserID     	uuid.UUID  `json:"user_id"`     // lounge_member.user_id
    Nickname   	string     `json:"nickname"`    
    Profile		string     `json:"profile"`
    RoleName   *string    `json:"role_name"`   // nullable: 역할이 없을 수도 있으므로 포인터
    RoleColor  *string    `json:"role_color"`  // nullable
}