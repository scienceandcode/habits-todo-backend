package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func GoogleAuthRoutes(group *gin.RouterGroup) {
	controller := &controller.GoogleAuthController{}

	group.GET("/setup", controller.StartGoogleAuthorization)
	group.GET("/code", controller.AuthorizationCode)
}
