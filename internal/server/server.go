package server

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/middleware"
	"github.com/scienceandcode/habits-todo-backend/internal/api/route"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

type HttpServer struct {
	healthController     *controller.HealthController
	googleAuthController *controller.GoogleAuthController
	userAuthController   *controller.UserAuthController
	userController       *controller.UserController
	adminController      *controller.AdminController
}

func (httpServer *HttpServer) registerRoutes(app *gin.Engine) {
	rootGroup := app.Group("/api")

	route.HealthRoutes(httpServer.healthController, rootGroup)
	route.GoogleAuthRoutes(httpServer.googleAuthController, rootGroup.Group("/google/auth"))
	route.UserAuthRoutes(httpServer.userAuthController, rootGroup.Group("/user/auth"))
	route.UserRoutes(httpServer.userController, rootGroup.Group("/user", middleware.AuthMiddleware()))
	route.AdminRoutes(httpServer.adminController, rootGroup.Group("/admin"))
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
	gac *controller.GoogleAuthController,
	uac *controller.UserAuthController,
	uc *controller.UserController,
	ac *controller.AdminController,
) *HttpServer {
	return &HttpServer{
		healthController:     hc,
		googleAuthController: gac,
		userAuthController:   uac,
		userController:       uc,
		adminController:      ac,
	}
}
