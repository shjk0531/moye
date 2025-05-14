package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserStudiesResponse는 사용자가 속한 스터디 목록을 반환하는 응답 DTO입니다.
type UserStudiesResponse struct {
	Studies []StudyListDTO `json:"studies"`
}

// study 목록 조회 시 사용되는 DTO
type StudyListDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProfileURL  string    `json:"profile_url"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserStudyDTO는 사용자의 스터디 정보를 표현하는 DTO입니다.
type UserStudyDTO struct {
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

// SimpleStudyListResponse는 스터디 목록을 조회하는 응답 DTO
type SimpleStudyListResponse struct {
	Studies []SimpleStudyDTO `json:"studies"`
}

// StudyIconResponse는 스터디 아이콘 목록을 조회하는 응답 DTO
type StudyIconResponse struct {
	Icons []StudyIconDTO `json:"icons"`
}