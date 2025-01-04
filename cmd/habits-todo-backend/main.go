package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"github.com/scienceandcode/habits-todo-backend/internal/server"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

func main() {
	log.Println("[HabitsTodo] Service Started")

	godotenv.Load()

	startHttpServer()

	common.WaitOsInterruption()
}

func startHttpServer() {
	httpServer := setupHttpServer()

	setupDatabase()

	log.Println("[HttpServer] Starting...")
	go httpServer.Run()
	log.Println("[HttpServer] Started")
}

func setupHttpServer() *server.HttpServer {
	healthController := controller.NewHealthController(service.NewHealthService())
	googleAuthController := controller.NewGoogleAuthController(service.NewGoogleAuthService())

	return server.NewHttpServer(
		healthController,
		googleAuthController,
	)
}

func setupDatabase() {
	log.Println("[Infrastructure] Connecting to database...")
	gormDbConnection := db.Init()
	log.Println("[Infrastructure] Database connected...")

	log.Println("[Infrastructure] Migrating pending models...")
	db.MigrateModels(gormDbConnection)
	log.Println("[Infrastructure] Models migrated...")
}
