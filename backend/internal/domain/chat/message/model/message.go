package model

import (
	"time"

	"github.com/kamva/mgm/v3"
)

// Message 모델은 채팅 메시지를 저장
type Message struct {
	mgm.DefaultModel `bson:",inline"`

    ChannelID    string    `bson:"channel_id"`
    UserID      string    `bson:"user_id"`
    Content   string    `bson:"content"`
    Type      string    `bson:"type"`
    CreatedAt time.Time `bson:"created_at"`
}
