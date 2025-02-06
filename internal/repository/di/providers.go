package di

import (
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}
