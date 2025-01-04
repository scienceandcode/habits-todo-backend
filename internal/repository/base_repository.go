package repository

import (
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{DB: db.GetConnection()}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	result := r.DB.Create(entity)
	return result.Error
}

func (r *BaseRepository[T]) FindByID(id any) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.DB.Find(&entities)
	return entities, result.Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	result := r.DB.Save(entity)
	return result.Error
}

func (r *BaseRepository[T]) Delete(id any) error {
	result := r.DB.Delete(new(T), id)
	return result.Error
}

func (r *BaseRepository[T]) Count() (int64, error) {
	var count int64
	result := r.DB.Model(new(T)).Count(&count)
	return count, result.Error
}

func (r *BaseRepository[T]) Paginate(page, pageSize int) ([]T, error) {
	var entities []T
	offset := (page - 1) * pageSize
	result := r.DB.Limit(pageSize).Offset(offset).Find(&entities)
	return entities, result.Error
}
