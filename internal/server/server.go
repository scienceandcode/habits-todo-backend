package server

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/route"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

type HttpServer struct {
	healthController *controller.HealthController
}

func (httpServer *HttpServer) registerRoutes(app *gin.Engine) {
	rootGroup := app.Group("/api")

	route.HealthRoutes(rootGroup)
}

func (httpServer *HttpServer) Run() {
	gin.SetMode(common.GetEnv("GIN_MODE"))
	gin.DisableConsoleColor()

	app := gin.Default()

	httpServer.registerRoutes(app)

	httpHandlerKey := common.GetEnv("HTTP_SERVER_HANDLER")
	Handle(httpHandlerKey, app)
}

func NewHttpServer(
	hc *controller.HealthController,
) *HttpServer {
	return &HttpServer{
		healthController: hc,
	}
}
