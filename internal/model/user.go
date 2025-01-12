package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	Name             string    `json:"name"`
	Email            string    `gorm:"uniqueIndex" json:"email"`
	Password         string    `json:"Password" validate:"min=6"`
	UserType         string    `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	EmailConfirmedAt time.Time `json:"email_confirmed_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	user.UpdatedAt = time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.UserType = string("USER")
	return nil
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
