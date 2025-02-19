package db

import (
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		&model.GoogleOAuthToken{},
		&model.User{},
	)
}