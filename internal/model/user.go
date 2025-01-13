package model

import (
	"time"

	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uint       `gorm:"primary_key" json:"id"`
	Name             string     `json:"name"`
	Email            string     `gorm:"uniqueIndex" json:"email"`
	Password         string     `json:"Password" validate:"min=8"`
	UserType         string     `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	EmailConfirmedAt *time.Time `json:"email_confirmed_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	user.EmailConfirmedAt = nil
	user.UserType = "USER"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	user.UpdatedAt = time.Now()
	return nil
}

func NewUserFromCreateUserRequestDTO(dto *dto.CreateUserRequestDTO) *User {
	return &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (user *User) ToUserDTO() *dto.UserDTO {
	return dto.NewUserDTO(user.ID, user.Name, user.Email, user.UserType, user.EmailConfirmedAt)
}
