package di

import (
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

func ProvideJWTService() *service.JWTService {
	return service.NewJWTService(common.GetEnv("JWT_SECRET_KEY"))
}
