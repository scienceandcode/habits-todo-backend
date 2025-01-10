package model

import (
	"time"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Email        string    `gorm:"uniqueIndex" json:"email"`
	Password     string    `json:"Password" validate:"min=6"`
	Token        *string   `json:"token"`
	UserType     *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	user.UpdatedAt = time.Now()
	user.Password = common.GenerateHMACUsingSHA256(user.Password, common.GetEnv("APP_ENCRYPTION_CYPHER_TEXT"))
	return nil
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
