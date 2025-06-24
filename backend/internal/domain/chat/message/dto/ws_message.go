package dto

import (
	"time"

	"github.com/shjk0531/moye/backend/internal/domain/chat/message/model"
)

type WsMessage struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	SenderID    string `json:"sender_id"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func FromModel(model *model.Message) *WsMessage {
	return &WsMessage{
		ID:        model.ID.Hex(),
		ChannelID: model.ChannelID,
		SenderID:    model.UserID,
		Content:   model.Content,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	}
}