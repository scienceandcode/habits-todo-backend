//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/scienceandcode/habits-todo-backend/internal/api/controller"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
	serviceDI "github.com/scienceandcode/habits-todo-backend/internal/api/service/di"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	repositoryDI "github.com/scienceandcode/habits-todo-backend/internal/repository/di"
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
		repositoryDI.UserRepositorySet,
		serviceDI.ProvideJWTService,
		db.GetConnection,
	)
	return &server.HttpServer{}
}
