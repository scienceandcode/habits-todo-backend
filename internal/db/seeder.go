package db

import (
	"log"

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

	result := db.Create(&adminUser)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	log.Println("[Seed] Admin user created")
}
