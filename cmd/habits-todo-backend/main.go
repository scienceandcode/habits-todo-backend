package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
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
	healthController := controller.NewHealthController(service.NewHealthService())

	httpServer := server.NewHttpServer(healthController)
	log.Println("[HttpServer] Starting...")
	go httpServer.Run()
	log.Println("[HttpServer] Started")
}
