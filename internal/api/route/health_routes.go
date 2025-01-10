package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func HealthRoutes(controller *controller.HealthController, group *gin.RouterGroup) {
	group.GET("/health", controller.GetHealth)
}
