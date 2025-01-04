package model

import (
	"time"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"gorm.io/gorm"
)

type GoogleOAuthToken struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
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
	token.AccessToken = common.DecryptAES(token.AccessToken)
	token.RefreshToken = common.DecryptAES(token.RefreshToken)

	return nil
}

func NewGoogleOAuthToken(accessToken, tokenType, refreshToken string, expiresIn int) *GoogleOAuthToken {
	return &GoogleOAuthToken{
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}
}
