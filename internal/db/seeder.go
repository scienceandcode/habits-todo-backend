package db

import (
	"github.com/scienceandcode/habits-todo-backend/internal/api/logger"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	adminUser := model.User{
		Name:     "Science & Code",
		Email:    common.GetEnv("ADMIN_USER_EMAIL"),
		Password: common.GetEnv("ADMIN_USER_PASSWORD"),
		UserType: "ADMIN",
	}

	var existingAdminUser model.User
	db.Where("email = ?", adminUser.Email).First(&existingAdminUser)

	if existingAdminUser.ID != 0 {
		logger.Info("[Seed] Admin user already exists")
		return
	}

	result := db.Create(&adminUser)

	if result.Error != nil {
		logger.Fatal(result.Error.Error())
	}

	logger.Info("[Seed] Admin user created")
}
