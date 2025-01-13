package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func UserAuthRoutes(controller *controller.UserAuthController, group *gin.RouterGroup) {
	group.POST("/register", controller.Register)
}
