package db

import (
	"fmt"
	"log"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	user := common.GetEnv("POSTGRES_USER")
	password := common.GetEnv("POSTGRES_PASSWORD")
	host := common.GetEnv("POSTGRES_HOST")
	port := common.GetEnv("POSTGRES_PORT")
	dbname := common.GetEnv("POSTGRES_NAME")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
