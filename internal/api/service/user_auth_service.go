package service

import (
	"log"

	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

type UserAuthService struct{}

func (service *UserAuthService) Register(dto *dto.CreateUserRequestDTO) (*dto.UserDTO, *errors.Error) {
	errorsList := service.validateCreateUserRequestDTO(dto)

	if errorsList != nil {
		return nil, errors.NewError("User registration failed.", errorsList)
	}

	userRepository := repository.NewRepository[model.User]()
	user := model.NewUserFromCreateUserRequestDTO(dto)

	repoErr := userRepository.Create(user)

	if repoErr != nil {
		log.Printf("Failed to create user: %v", repoErr.Error())
		return nil, errors.NewError(repoErr.Error(), nil)
	}

	return user.ToUserDTO(), nil
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

func (service *UserAuthService) validateUserEmail(email string) *errors.FieldError {
	if !common.ValidateEmail(email) {
		return errors.NewFieldError("email", "Please provide a valid email address.")
	}

	userRepository := repository.NewRepository[model.User]()
	user, _ := userRepository.FindOneBy(map[string]interface{}{"email": email})

	if user != nil {
		return errors.NewFieldError("email", "Email address is already in use.")
	}

	return nil
}

func NewUserAuthService() *UserAuthService {
	return &UserAuthService{}
}
