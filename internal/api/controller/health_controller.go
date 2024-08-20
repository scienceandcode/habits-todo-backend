package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type HealthController struct {
	service *service.HealthService
}

func (controller *HealthController) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, controller.service.GetHealthMessage())
}

func NewHealthController(service *service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}
