package dto

type CreateChannelRequest struct {
	Name string `json:"name" binding:"required"`
	Position int `json:"position" binding:"required"`
}
