package controller

import (
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type AuthController struct {
	service service.Service
}

func NewAuthController(service service.Service) *AuthController {
	return &AuthController{service: service}
}

