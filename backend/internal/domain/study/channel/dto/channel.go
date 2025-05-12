package dto

import "github.com/google/uuid"


type ChannelDTO struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}


type TreeItemDTO struct {
	Type string `json:"type"`
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Channels []ChannelDTO `json:"channels,omitempty"`
}
