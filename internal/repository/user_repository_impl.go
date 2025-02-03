package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*Repository[model.User]
}

func (r *UserRepositoryImpl) FindByID(id uint) (*model.User, error) {
	return r.Repository.FindByID(id)
}

func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.Repository.Create(user)
}

func (r *UserRepositoryImpl) FindOneBy(condition map[string]interface{}) (*model.User, error) {
	return r.Repository.FindOneBy(condition)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Repository: &Repository[model.User]{DB: db},
	}
}
