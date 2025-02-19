package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type AdminController struct {
	service *service.AdminService
}

func (controller *AdminController) MigrateModels(c *gin.Context) {
	controller.service.MigrateModels(c)
	api.ResponseSuccess(c, http.StatusCreated, gin.H{"message": "Migration and seeding completed"})
}

func NewAdminController(service *service.AdminService) *AdminController {
	return &AdminController{
		service: service,
	}
}
