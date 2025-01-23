package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		userID, err := common.ValidateJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
