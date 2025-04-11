package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

// VideoChatSession 모델은 화상채팅 세션 정보를 저장
type VideoChatSession struct {
	model.BaseModel
	ChatRoomID *uuid.UUID     	`gorm:"not null" json:"chatroom_id"`        // 채팅방 외래키 (화상채팅 전용 또는 그룹채팅 연계)
	SessionURL string   		`gorm:"not null" json:"session_url"`	  // 접속 URL 또는 세션 정보
	CreatedBy  *uuid.UUID     	`gorm:"not null" json:"created_by"` // 세션을 시작한 사용자 (User 외래키)
}
