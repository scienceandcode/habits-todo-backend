package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/api/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			api.ResponseUnauthorized(c, errors.NewSimpleError("Authorization header is required"))
			c.Abort()
			return
		}

		jwtService := service.NewJWTService()

		userID, err := jwtService.ValidateJWT(authHeader)
		if err != nil {
			api.ResponseUnauthorized(c, errors.NewSimpleError("Invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
