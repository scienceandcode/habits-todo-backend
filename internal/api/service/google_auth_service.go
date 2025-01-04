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
	state := common.EncryptAES(common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY"))
	nonce := string(time.Now().UnixNano())
	scope := "openid%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar"

	url := "https://accounts.google.com/o/oauth2/v2/auth?accessType=offline&response_type=code&client_id=%s&redirect_uri=%s&state=%s&nonce=%s&scope=%s"

	return fmt.Sprintf(url, clientId, redirectUri, state, nonce, scope)
}

func (service *GoogleAuthService) ExchangeCodeForToken(code, state string) error {
	//clientId := common.GetEnv("GOOGLE_CLOUD_CLIENT_ID")
	//clientSecret := common.GetEnv("GOOGLE_CLOUD_CLIENT_SECRET")
	//redirectUri := common.GetEnv("GOOGLE_CLOUD_REDIRECT_URI")

	decryptedState, _ := common.DecryptAES(state)
	if common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY") != decryptedState {
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
