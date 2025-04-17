package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=20"`
	Nickname string `json:"nickname" binding:"required,min=2,max=10"`
}