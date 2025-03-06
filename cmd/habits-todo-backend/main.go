package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"github.com/scienceandcode/habits-todo-backend/internal/server/di"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"github.com/scienceandcode/habits-todo-backend/pkg/environment"
)

func main() {
	log.Println("[HabitsTodo] Service Started")

	godotenv.Load()

	setupDatabase()
	startHttpServer()

	common.WaitOsInterruption()
}

func startHttpServer() {
	httpServer := di.InitializeHttpServer()

	log.Println("[HttpServer] Starting...")
	go httpServer.Run()
	log.Println("[HttpServer] Started")
}

func setupDatabase() {
	log.Println("[Infrastructure] Connecting to database...")
	gormDbConnection := db.Init()
	if environment.IsDevelopment() {
		db.MigrateModels(gormDbConnection)
		db.SeedAdminUser(gormDbConnection)
	}
	log.Println("[Infrastructure] Database connected...")
}
