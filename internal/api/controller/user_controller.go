package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

type UserController struct {
	service *service.UserService
}

func (controller *UserController) Profile(c *gin.Context) {
	user, err := controller.service.Profile(c)

	if err != nil {
		api.ResponseNotFound(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusCreated, user)
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
