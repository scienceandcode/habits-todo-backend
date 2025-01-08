package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{DB: db.GetConnection()}
}

func (r *Repository[T]) Create(entity *T) error {
	result := r.DB.Create(entity)
	return result.Error
}

func (r *Repository[T]) FindByID(id any) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *Repository[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.DB.Find(&entities)
	return entities, result.Error
}

func (r *Repository[T]) Update(entity *T) error {
	result := r.DB.Save(entity)
	return result.Error
}

func (r *Repository[T]) Delete(id any) error {
	result := r.DB.Delete(new(T), id)
	return result.Error
}

func (r *Repository[T]) Count() (int64, error) {
	var count int64
	result := r.DB.Model(new(T)).Count(&count)
	return count, result.Error
}

func (r *Repository[T]) Paginate(page, pageSize int) ([]T, error) {
	var entities []T
	offset := (page - 1) * pageSize
	result := r.DB.Limit(pageSize).Offset(offset).Find(&entities)
	return entities, result.Error
}

func (r *Repository[T]) FindOneBy(condition map[string]interface{}) (*T, error) {
	var entity T
	result := r.DB.Where(condition).First(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}
