package db

import (
	"github.com/scienceandcode/habits-todo-backend/internal/api/logger"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	logger.Info("[Infrastructure] Migrating pending models...")
	db.AutoMigrate(
		&model.GoogleOAuthToken{},
		&model.User{},
	)
	logger.Info("[Infrastructure] Models migrated...")
}
