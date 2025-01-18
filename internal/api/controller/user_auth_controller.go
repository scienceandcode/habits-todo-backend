package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/api/dto"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type UserAuthController struct {
	service *service.UserAuthService
}

func (controller *UserAuthController) Register(c *gin.Context) {
	var createUserRequestDTO *dto.CreateUserRequestDTO
	api.ParseRequest(c, &createUserRequestDTO)

	user, err := controller.service.Register(createUserRequestDTO)

	if err != nil {
		api.ResponseBadRequest(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusCreated, user)
}

func NewUserAuthController(service *service.UserAuthService) *UserAuthController {
	return &UserAuthController{
		service: service,
	}
}
