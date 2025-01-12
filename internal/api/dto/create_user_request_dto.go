package dto

type CreateUserRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

func NewCreateUserRequestDTO(name, email, password string) *CreateUserRequestDTO {
	return &CreateUserRequestDTO{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
