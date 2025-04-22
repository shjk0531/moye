package dto

import (
	"github.com/google/uuid"
)

// ──────────────────────────────────────────
// 1) 최상위 응답: 스터디 채널 리스트
// ──────────────────────────────────────────
type StudyChannelsResponse struct {
    Items []StudyChannelItem `json:"items"`
}

// ──────────────────────────────────────────
// 2) 리스트 항목: 단독 채널 혹은 그룹
// ──────────────────────────────────────────
type StudyChannelItem struct {
    ItemType string           `json:"item_type"`           // "channel" or "group"
    Channel  *ChannelDTO      `json:"channel,omitempty"`   // ItemType=="channel"일 때
    Group    *ChannelGroupDTO `json:"group,omitempty"`     // ItemType=="group"일 때
}

// ──────────────────────────────────────────
// 3) 단독 채널 DTO
// ──────────────────────────────────────────
type ChannelDTO struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
    // 필요 시 추가 메타(ex. unread_count, last_message 등)도 여기에 넣을 수 있습니다.
}

// ──────────────────────────────────────────
// 4) 채널 그룹 DTO
// ──────────────────────────────────────────
type ChannelGroupDTO struct {
    ID       uuid.UUID    `json:"id"`
    Name     string       `json:"name"`
    Channels []ChannelDTO `json:"channels"`  // 그룹 내부에서 position 순으로 정렬된 채널 목록
}
