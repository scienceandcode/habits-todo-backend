package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
)

func ParseRequest(c *gin.Context, dto any) error {
	if err := c.BindJSON(&dto); err != nil {
		return err
	}
	return nil
}

func ResponseBadRequest(c *gin.Context, err *errors.Error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func ResponseSuccess(c *gin.Context, httpStatusCode int, data any) {
	c.JSON(httpStatusCode, data)
}

func ResponseUnauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err})
}

func ResponseNotFound(c *gin.Context, err *errors.Error) {
	c.JSON(http.StatusNotFound, gin.H{"error": err})
}
