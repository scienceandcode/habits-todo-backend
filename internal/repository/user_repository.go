package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/model"
)

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	FindOneBy(condition map[string]interface{}) (*model.User, error)
}
