package route

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
)

func HealthRoutes(group *gin.RouterGroup) {
	controller := &controller.HealthController{}

	group.GET("/health", controller.GetHealth)
}
