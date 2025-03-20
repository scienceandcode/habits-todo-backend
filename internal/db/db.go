package db

import (
	"fmt"
	"log"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gormDbConnection *gorm.DB

func GetConnection() *gorm.DB {
	return gormDbConnection
}

func InitTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	gormDbConnection = db
}

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

	gormDbConnection = db

	return db
}