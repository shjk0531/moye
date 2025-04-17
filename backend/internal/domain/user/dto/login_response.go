package dto

import "github.com/google/uuid"

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        User   `json:"user"`	
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Nickname string    `json:"nickname"`
	Profile  *string    `json:"profile"`
}
