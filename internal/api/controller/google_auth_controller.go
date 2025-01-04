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

func (controller *GoogleAuthController) AuthorizationCode(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" || state == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code and state are required"})
		return
	}

	err := controller.service.ExchangeCodeForToken(code, state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func NewGoogleAuthController(service *service.GoogleAuthService) *GoogleAuthController {
	return &GoogleAuthController{
		service: service,
	}
}
