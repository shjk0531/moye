package model

import (
	"time"

	"github.com/kamva/mgm/v3"
)

// Message 모델은 채팅 메시지를 저장
type Message struct {
	mgm.DefaultModel `bson:",inline"`

    ChannelID    string    `bson:"channel_id"`     // ChatRoom.ID
    UserID      string    `bson:"user_id"`
    Content   string    `bson:"content"`
    Type      string    `bson:"type"`        // text, image, file 등
    CreatedAt time.Time `bson:"created_at"`
}
