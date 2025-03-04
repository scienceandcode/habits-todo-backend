package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
)

type AdminController struct{}

func (controller *AdminController) MigrateModels(c *gin.Context) {
	db.MigrateModels(db.GetConnection())
	api.ResponseSuccess(c, http.StatusCreated, gin.H{"message": "Model migrations completed"})
}

func NewAdminController() *AdminController {
	return &AdminController{}
}
