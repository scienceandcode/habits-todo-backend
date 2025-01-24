//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
	"github.com/scienceandcode/habits-todo-backend/internal/server"
)

func InitializeHttpServer() *server.HttpServer {
	wire.Build(
		server.NewHttpServer,
		controller.NewHealthController,
		controller.NewGoogleAuthController,
		controller.NewUserAuthController,
		controller.NewUserController,
		service.NewHealthService,
		service.NewGoogleAuthService,
		service.NewUserAuthService,
		service.NewUserService,
	)
	return &server.HttpServer{}
}
