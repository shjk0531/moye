package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;index" json:"updated_by"`
	Version   int        `gorm:"version;default:1" json:"version"`
}

// BeforeCreate sets the created/updated by fields and generates a UUID if not already set.
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	// 컨텍스트가 없는 경우 바로 리턴
	if tx.Statement.Context == nil {
		return nil
	}

	// UUID 자동 생성 (값이 없는 경우)
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}

	// 컨텍스트에서 사용자 ID 가져오기
	if userID, ok := middleware.GetUserID(tx.Statement.Context); ok {
		b.CreatedBy = &userID
		b.UpdatedBy = &userID
	} else {
		log.Println("Warning: no user ID found in context during BeforeCreate")
	}
	return nil
}

// BeforeUpdate sets the updated by field.
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	// 컨텍스트가 없는 경우 바로 리턴
	if tx.Statement.Context == nil {
		return nil
	}

	// 컨텍스트에서 사용자 ID 가져오기
	if userID, ok := middleware.GetUserID(tx.Statement.Context); ok {
		b.UpdatedBy = &userID
	} else {
		log.Println("Warning: no user ID found in context during BeforeUpdate")
	}
	return nil
}

// AfterCreate is a hook for post-create processing (e.g., event publishing, cache update).
func (b *BaseModel) AfterCreate(tx *gorm.DB) error {
	log.Printf("Model created with ID: %s", b.ID)
	// 추가 후처리 작업을 이곳에 구현할 수 있습니다.
	return nil
}

// AfterUpdate is a hook for post-update processing (e.g., event publishing, cache update).
func (b *BaseModel) AfterUpdate(tx *gorm.DB) error {
	log.Printf("Model updated with ID: %s, new version: %d", b.ID, b.Version)
	// 추가 후처리 작업을 이곳에 구현할 수 있습니다.
	return nil
}
