package dto

type TokenResponseDTO struct {
	Token string `json:"token"`
}

func NewTokenResponseDTO(token string) *TokenResponseDTO {
	return &TokenResponseDTO{Token: token}
}
