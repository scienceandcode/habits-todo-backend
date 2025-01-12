package dto

type CreateUserRequestDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserRequestDTO(name, email, password string) *CreateUserRequestDTO {
	return &CreateUserRequestDTO{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
