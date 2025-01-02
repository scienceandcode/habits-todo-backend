package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type GoogleAuthController struct {
	service *service.GoogleAuthService
}

func (controller *GoogleAuthController) StartGoogleAuthorization(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, controller.service.BuildGoogleAuthURL())
}

func NewGoogleAuthController(service *service.GoogleAuthService) *GoogleAuthController {
	return &GoogleAuthController{
		service: service,
	}
}
