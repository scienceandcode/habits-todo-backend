package integration

type GoogleOAuthTokenResponseDTO struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func NewGoogleOAuthTokenResponseDTO() *GoogleOAuthTokenResponseDTO {
	return &GoogleOAuthTokenResponseDTO{}
}
