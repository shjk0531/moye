package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserLoungesResponse는 사용자가 속한 스터디 목록을 반환하는 응답 DTO입니다.
type UserLoungesResponse struct {
	Lounges []LoungeListDTO `json:"lounges"`
}

// lounge 목록 조회 시 사용되는 DTO
type LoungeListDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProfileURL  string    `json:"profile_url"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserLoungeDTO는 사용자의 스터디 정보를 표현하는 DTO입니다.
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

// UserRoleDTO는 해당 스터디에서 사용자의 역할 정보를 표현하는 DTO입니다.
type UserRoleDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ColorHex  string    `json:"color_hex"`
	RoleFlags int64     `json:"role_flags"`
}

// SimpleLoungeListResponse는 스터디 목록을 조회하는 응답 DTO
type SimpleLoungeListResponse struct {
	Lounges []SimpleLoungeDTO `json:"lounges"`
}

// LoungeIconResponse는 스터디 아이콘 목록을 조회하는 응답 DTO
type LoungeIconResponse struct {
	Icons []LoungeIconDTO `json:"icons"`
}