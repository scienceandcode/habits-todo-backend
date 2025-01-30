package service

import (
	"log"

	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthService struct {
	UserRepo   repository.UserRepository
	JwtService *JWTService
}

func (service *UserAuthService) Register(dto *dto.CreateUserRequestDTO) (*dto.UserDTO, *errors.Error) {
	errorsList := service.validateCreateUserRequestDTO(dto)

	if errorsList != nil {
		return nil, errors.NewError("User registration failed.", errorsList)
	}

	user := model.NewUserFromCreateUserRequestDTO(dto)
	repoErr := service.UserRepo.Create(user)

	if repoErr != nil {
		log.Printf("Failed to create user: %v", repoErr.Error())
		return nil, errors.NewError(repoErr.Error(), nil)
	}

	return user.ToUserDTO(), nil
}

func (service *UserAuthService) Login(loginRequestDTO *dto.LoginRequestDTO) (*dto.TokenResponseDTO, *errors.Error) {
	if errorsList := service.validateLoginRequestDTO(loginRequestDTO); errorsList != nil {
		return nil, errors.NewError("Invalid login request.", errorsList)
	}

	user, _ := service.UserRepo.FindOneBy(map[string]interface{}{"email": loginRequestDTO.Email})

	if credErr := service.validateUserCredentials(user, loginRequestDTO.Password); credErr != nil {
		return nil, errors.NewError("Invalid credentials.", []*errors.FieldError{credErr})
	}

	token, err := service.JwtService.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.NewError("Error generating token.", nil)
	}

	response := dto.NewTokenResponseDTO(token)

	return response, nil
}

func (service *UserAuthService) validateCreateUserRequestDTO(dto *dto.CreateUserRequestDTO) []*errors.FieldError {
	var errorsList []*errors.FieldError

	emailError := service.validateUserEmail(dto.Email)
	if emailError != nil {
		errorsList = append(errorsList, emailError)
	}

	if len(dto.Password) < 8 {
		errorsList = append(errorsList, errors.NewFieldError("password", "Password must be at least 8 characters long."))
	}

	if dto.Name == "" {
		errorsList = append(errorsList, errors.NewFieldError("name", "Name is required."))
	}

	if len(errorsList) > 0 {
		return errorsList
	}

	return nil
}

func (service *UserAuthService) validateLoginRequestDTO(dto *dto.LoginRequestDTO) []*errors.FieldError {
	var errorsList []*errors.FieldError

	if dto.Email == "" {
		errorsList = append(errorsList, errors.NewFieldError("email", "Email is required."))
	} else if !common.ValidateEmail(dto.Email) {
		errorsList = append(errorsList, errors.NewFieldError("email", "Please provide a valid email address."))
	}

	if dto.Password == "" {
		errorsList = append(errorsList, errors.NewFieldError("password", "Password is required."))
	}

	if len(errorsList) > 0 {
		return errorsList
	}

	return nil
}

func (service *UserAuthService) validateUserCredentials(user *model.User, password string) *errors.FieldError {
	if user == nil {
		return errors.NewFieldError("email", "Email not found.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.NewFieldError("password", "Incorrect password.")
	}

	return nil
}

func (service *UserAuthService) validateUserEmail(email string) *errors.FieldError {
	if !common.ValidateEmail(email) {
		return errors.NewFieldError("email", "Please provide a valid email address.")
	}

	user, _ := service.UserRepo.FindOneBy(map[string]interface{}{"email": email})

	if user != nil {
		return errors.NewFieldError("email", "Email address is already in use.")
	}

	return nil
}

func NewUserAuthService(userRepo repository.UserRepository, jwtService *JWTService) *UserAuthService {
	return &UserAuthService{UserRepo: userRepo, JwtService: jwtService}
}
