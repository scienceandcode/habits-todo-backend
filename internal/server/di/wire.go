//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"github.com/scienceandcode/habits-todo-backend/internal/server"
)

func InitializeHttpServer() *server.HttpServer {
	wire.Build(
		server.NewHttpServer,
		controller.NewAdminController,
		controller.NewGoogleAuthController,
		controller.NewHealthController,
		controller.NewUserAuthController,
		controller.NewUserController,
		service.NewAdminService,
		service.NewGoogleAuthService,
		service.NewHealthService,
		service.NewJWTService,
		service.NewUserAuthService,
		service.NewUserService,
		repository.NewUserRepository,
		db.GetConnection,
	)
	return &server.HttpServer{}
}
