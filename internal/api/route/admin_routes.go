package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func AdminRoutes(controller *controller.AdminController, adminGroup *gin.RouterGroup) {
	adminGroup.POST("/migrate", controller.MigrateModels)
}
