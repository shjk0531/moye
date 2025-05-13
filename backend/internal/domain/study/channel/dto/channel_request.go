package dto

import "github.com/google/uuid"

type CreateChannelRequest struct {
	Name string `json:"name" binding:"required"`
	Position int `json:"position" binding:"required"`
}

// 채널 순서 재배치 요청
type ReorderChannelsRequest struct {
    Items []ReorderChannelItem `json:"items" binding:"required,min=1"`
}

type ReorderChannelItem struct {
    ItemType string    `json:"item_type" binding:"required,oneof=channel group"`
    ItemID   uuid.UUID `json:"item_id" binding:"required"`
}