package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/api/logger"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"github.com/scienceandcode/habits-todo-backend/pkg/integration"
)

type GoogleAuthService struct{}

func (*GoogleAuthService) BuildGoogleAuthURL() string {
	clientId := common.GetEnv("GOOGLE_CLOUD_CLIENT_ID")
	redirectUri := common.GetEnv("GOOGLE_CLOUD_REDIRECT_URI")
	state := common.EncryptAES(common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY"))
	nonce := string(time.Now().UnixNano())
	scope := "openid%20profile%20email%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar"

	url := "https://accounts.google.com/o/oauth2/v2/auth?access_type=offline&response_type=code&client_id=%s&redirect_uri=%s&state=%s&nonce=%s&scope=%s"

	return fmt.Sprintf(url, clientId, redirectUri, state, nonce, scope)
}

func (service *GoogleAuthService) ExchangeCodeForToken(code, state string) *errors.Error {

	validationErrors := service.validateCodeAndState(code, state)

	if validationErrors != nil {
		return validationErrors
	}

	decryptedState, _ := common.DecryptAES(state)
	defaultErrorMessage := "Error while exchanging code for token"

	if common.GetEnv("GOOGLE_CLOUD_AUTH_STATE_SECRET_KEY") != decryptedState {
		return errors.NewError(defaultErrorMessage, []*errors.FieldError{errors.NewFieldError("state", "Invalid state")})
	}

	httpClient := &http.Client{}
	res, err := httpClient.Post("https://oauth2.googleapis.com/token", "application/x-www-form-urlencoded", strings.NewReader(service.buildTokenRequestFormData(code).Encode()))

	if err != nil {
		logger.Error(fmt.Sprintf("%s: %s", defaultErrorMessage, err.Error()))
		return errors.NewError(defaultErrorMessage, []*errors.FieldError{errors.NewFieldError("statusCode", "Auth request failed")})
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		logger.Error(fmt.Sprintf("%s: %s", defaultErrorMessage, string(body)))
		return errors.NewError(defaultErrorMessage, []*errors.FieldError{errors.NewFieldError("statusCode", "Auth request failed")})
	}

	tokenResponseDTO := integration.NewGoogleOAuthTokenResponseDTO()
	err = json.NewDecoder(res.Body).Decode(tokenResponseDTO)

	if err != nil {
		logger.Error(fmt.Sprintf("%s: %s", defaultErrorMessage, err.Error()))
		return errors.NewError(defaultErrorMessage, []*errors.FieldError{errors.NewFieldError("response", "Invalid auth response body")})
	}

	err = service.saveToken(tokenResponseDTO)

	if err != nil {
		logger.Error(fmt.Sprintf("%s: %s", defaultErrorMessage, err.Error()))
		return errors.NewError(defaultErrorMessage, []*errors.FieldError{errors.NewFieldError("token", "Error while saving token")})
	}

	return nil
}

func (*GoogleAuthService) validateCodeAndState(code, state string) *errors.Error {
	dataValidationErrors := []*errors.FieldError{}

	if code == "" {
		dataValidationErrors = append(dataValidationErrors, errors.NewFieldError("code", "Code is required"))
	}

	if state == "" {
		dataValidationErrors = append(dataValidationErrors, errors.NewFieldError("state", "State is required"))
	}

	if len(dataValidationErrors) > 0 {
		return errors.NewError("Invalid request data", dataValidationErrors)
	}

	return nil
}

func (service *GoogleAuthService) saveToken(tokenResponseDTO *integration.GoogleOAuthTokenResponseDTO) error {
	userEmail, err := service.getUserEmailFromAccessToken(tokenResponseDTO.AccessToken)

	if err != nil {
		return fmt.Errorf("error while getting user email: %v", err)
	}

	tokenRepository := repository.NewRepository[model.GoogleOAuthToken]()
	token := model.NewGoogleOAuthTokenFromDTO(tokenResponseDTO)
	token.UserEmail = userEmail

	existingToken, _ := tokenRepository.FindOneBy(map[string]interface{}{"user_email": userEmail})

	if existingToken != nil {
		token.ID = existingToken.ID
		tokenRepository.Update(token)
		return nil
	}

	tokenRepository.Create(token)

	return nil
}

func (service *GoogleAuthService) getUserEmailFromAccessToken(accessToken string) (string, error) {
	httpClient := &http.Client{}

	userInfoUrl, err := service.getUserInfoUrl()

	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("GET", userInfoUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := httpClient.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		return "", err
	}

	defer res.Body.Close()

	userInfoDTO := integration.NewGoogleUserInfoDTO()
	err = json.NewDecoder(res.Body).Decode(userInfoDTO)

	if err != nil {
		return "", err
	}

	return userInfoDTO.Email, nil
}

func (*GoogleAuthService) getUserInfoUrl() (string, error) {
	httpClient := &http.Client{}
	res, err := httpClient.Get("https://accounts.google.com/.well-known/openid-configuration")

	if err != nil || res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error while getting openid connect configuration: %v", err)
	}

	defer res.Body.Close()

	googleOpenIdConnectConfigurationDTO := integration.NewGoogleOpenIdConnectConfigurationDTO()
	err = json.NewDecoder(res.Body).Decode(googleOpenIdConnectConfigurationDTO)

	if err != nil {
		return "", fmt.Errorf("error while decoding openid connect config response body: %v", err)
	}

	return googleOpenIdConnectConfigurationDTO.UserInfoEndpoint, nil
}

func (*GoogleAuthService) buildTokenRequestFormData(code string) url.Values {
	clientId := common.GetEnv("GOOGLE_CLOUD_CLIENT_ID")
	clientSecret := common.GetEnv("GOOGLE_CLOUD_CLIENT_SECRET")
	redirectUri := common.GetEnv("GOOGLE_CLOUD_REDIRECT_URI")

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	return data
}

func NewGoogleAuthService() *GoogleAuthService {
	return &GoogleAuthService{}
}
