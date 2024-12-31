package controller

import (
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type GoogleAuthController struct {
	service *service.GoogleAuthService
}

func NewGoogleAuthController(service *service.GoogleAuthService) *GoogleAuthController {
	return &GoogleAuthController{
		service: service,
	}
}
