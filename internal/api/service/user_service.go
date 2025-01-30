package service

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/model"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
)

type UserService struct{}

func (service *UserService) Profile(c *gin.Context) (*dto.UserDTO, *errors.Error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, errors.NewError("User ID not found in context", nil)
	}

	id, ok := userID.(uint)
	if !ok {
		return nil, errors.NewError("Invalid user ID type in context", nil)
	}

	userRepository := repository.NewRepository[model.User]()
	user, err := userRepository.FindByID(id)
	if err != nil || user == nil {
		return nil, errors.NewError("User not found", nil)
	}

	return user.ToUserDTO(), nil
}

func NewUserService() *UserService {
	return &UserService{}
}
