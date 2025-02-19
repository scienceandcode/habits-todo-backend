package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
)

type AdminService struct{}

func (s *AdminService) MigrateModels(c *gin.Context) {
	dbConn := db.GetConnection()
	db.MigrateModels(dbConn)
	db.SeedAdminUser(dbConn)
	c.JSON(http.StatusOK, gin.H{"message": "Migration and seeding completed"})
}

func NewAdminService() *AdminService {
	return &AdminService{}
}
