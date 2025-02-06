package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	*Repository[model.User]
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	return r.Repository.FindByID(id)
}

func (r *userRepository) Create(user *model.User) error {
	return r.Repository.Create(user)
}

func (r *userRepository) FindOneBy(condition map[string]interface{}) (*model.User, error) {
	return r.Repository.FindOneBy(condition)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Repository: &Repository[model.User]{DB: db},
	}
}
