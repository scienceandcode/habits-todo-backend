package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func UserRoutes(controller *controller.UserController, group *gin.RouterGroup) {
	group.GET("/profile", controller.Profile)
}
