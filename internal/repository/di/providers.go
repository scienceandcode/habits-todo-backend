package di

import (
	"github.com/google/wire"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

var UserRepositorySet = wire.NewSet(
	ProvideUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)
