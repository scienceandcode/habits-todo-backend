package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository[model.User]{DB: db},
	}
}
