package service

import (
	"fmt"
	"time"

	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

type GoogleAuthService struct{}

func (*GoogleAuthService) BuildGoogleAuthURL() string {
	clientId := common.GetEnv("GOOGLE_CLOUD_CLIENT_ID")
	redirectUri := common.GetEnv("GOOGLE_CLOUD_REDIRECT_URI")
	scope := "openid%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar"
	nonce := string(time.Now().UnixNano())
	state := common.GenerateHMACUsingSHA256(clientId, common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY"))

	url := "https://accounts.google.com/o/oauth2/v2/auth?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&nonce=%s&state=%s"

	return fmt.Sprintf(url, clientId, redirectUri, scope, nonce, state)
}

func (service *GoogleAuthService) ExchangeCodeForToken(code, state string) error {
	clientId := common.GetEnv("GOOGLE_CLOUD_CLIENT_ID")
	//clientSecret := common.GetEnv("GOOGLE_CLOUD_CLIENT_SECRET")
	//redirectUri := common.GetEnv("GOOGLE_CLOUD_REDIRECT_URI")

	if !common.VerifyHMAC(clientId, common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY"), state) {
		return fmt.Errorf("invalid state")
	}

	// TODO: exchange code for token
	// TODO: handle token exchange error

	// TODO: save token to database instantiating repository
	// TODO: handle database save error

	return nil
}

func NewGoogleAuthService() *GoogleAuthService {
	return &GoogleAuthService{}
}
