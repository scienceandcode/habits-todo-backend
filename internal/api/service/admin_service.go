package service

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
)

type AdminService struct{}

func (s *AdminService) MigrateModels(c *gin.Context) {
	dbConn := db.GetConnection()
	db.MigrateModels(dbConn)
}

func NewAdminService() *AdminService {
	return &AdminService{}
}
