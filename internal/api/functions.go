package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseRequest(c *gin.Context, dto any) error {
	if err := c.BindJSON(&dto); err != nil {
		return err
	}
	return nil
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func ResponseSuccess(c *gin.Context, httpStatusCode int, data any) {
	c.JSON(httpStatusCode, data)
}
