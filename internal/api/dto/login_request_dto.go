package dto

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
