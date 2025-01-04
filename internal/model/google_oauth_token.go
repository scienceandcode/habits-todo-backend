package model

import (
	"time"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"github.com/scienceandcode/habits-todo-backend/pkg/integration"
	"gorm.io/gorm"
)

type GoogleOAuthToken struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
	Scope        string    `json:"scope"`
	CreatedAt    time.Time `json:"created_at"`
}

func (token *GoogleOAuthToken) BeforeCreate(tx *gorm.DB) error {
	token.CreatedAt = time.Now()
	return nil
}

func (token *GoogleOAuthToken) BeforeSave(tx *gorm.DB) error {
	token.AccessToken = common.EncryptAES(token.AccessToken)
	token.RefreshToken = common.EncryptAES(token.RefreshToken)

	return nil
}

func (token *GoogleOAuthToken) AfterFind(tx *gorm.DB) error {
	token.AccessToken, _ = common.DecryptAES(token.AccessToken)
	token.RefreshToken, _ = common.DecryptAES(token.RefreshToken)

	return nil
}

func NewGoogleOAuthToken(accessToken, tokenType, refreshToken, scope string, expiresIn int) *GoogleOAuthToken {
	return &GoogleOAuthToken{
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		Scope:        scope,
	}
}

func NewGoogleOAuthTokenFromDTO(dto *integration.GoogleOAuthTokenResponseDTO) *GoogleOAuthToken {
	return &GoogleOAuthToken{
		AccessToken:  dto.AccessToken,
		TokenType:    dto.TokenType,
		RefreshToken: dto.RefreshToken,
		ExpiresIn:    dto.ExpiresIn,
		Scope:        dto.Scope,
	}
}
