package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"github.com/scienceandcode/habits-todo-backend/internal/server"
	"github.com/scienceandcode/habits-todo-backend/internal/server/di"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

func main() {
	log.Println("[HabitsTodo] Service Started")

	godotenv.Load()

	setupDatabase()
	startHttpServer()

	common.WaitOsInterruption()
}

func startHttpServer() {
	httpServer := setupHttpServer()

	log.Println("[HttpServer] Starting...")
	go httpServer.Run()
	log.Println("[HttpServer] Started")
}

func setupHttpServer() *server.HttpServer {
	return di.InitializeHttpServer()
}

func setupDatabase() {
	log.Println("[Infrastructure] Connecting to database...")
	gormDbConnection := db.Init()
	log.Println("[Infrastructure] Database connected...")

	log.Println("[Infrastructure] Migrating pending models...")
	db.MigrateModels(gormDbConnection)
	log.Println("[Infrastructure] Models migrated...")
}
