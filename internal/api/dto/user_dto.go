package dto

import "time"

type UserDTO struct {
	ID               uint       `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	UserType         string     `json:"user_type"`
	EmailConfirmedAt *time.Time `json:"email_confirmed_at"`
}

func NewUserDTO(id uint, name string, email string, userType string, emailConfirmedAt *time.Time) *UserDTO {
	return &UserDTO{ID: id, Name: name, Email: email, UserType: userType, EmailConfirmedAt: emailConfirmedAt}
}
