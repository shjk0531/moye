package dto

import (
	"github.com/google/uuid"
)

// ──────────────────────────────────────────
// 1) 최상위 응답: 라운지 채널 리스트
// ──────────────────────────────────────────
type LoungeChannelsResponse struct {
    Items []LoungeChannelItem `json:"items"`
}

// ──────────────────────────────────────────
// 2) 리스트 항목: 단독 채널 혹은 그룹
// ──────────────────────────────────────────
type LoungeChannelItem struct {
    ItemType string           `json:"item_type"`           // "channel" or "group"
    Channel  *ChannelDTO      `json:"channel,omitempty"`   // ItemType=="channel"일 때
    Group    *ChannelGroupDTO `json:"group,omitempty"`     // ItemType=="group"일 때
}


// ──────────────────────────────────────────
// 4) 채널 그룹 DTO
// ──────────────────────────────────────────
type ChannelGroupDTO struct {
    ID       uuid.UUID    `json:"id"`
    Name     string       `json:"name"`
    Channels []ChannelDTO `json:"channels"`  // 그룹 내부에서 position 순으로 정렬된 채널 목록
}
