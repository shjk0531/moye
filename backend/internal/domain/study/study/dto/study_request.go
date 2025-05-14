package dto

type CreateStudyRequest struct {
	Name string `json:"name" binding:"required"`
	ProfileURL string `json:"profile_url" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content string `json:"content" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}
