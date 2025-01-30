package service

import (
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) Profile(c *gin.Context) (*dto.UserDTO, *errors.Error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, errors.NewError("User ID not found in context", nil)
	}

	id, ok := userID.(uint)
	if !ok {
		return nil, errors.NewError("Invalid user ID type in context", nil)
	}

	user, err := s.UserRepo.FindByID(id)
	if err != nil || user == nil {
		return nil, errors.NewError("User not found", nil)
	}

	return user.ToUserDTO(), nil
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}
