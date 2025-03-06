package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/scienceandcode/habits-todo-backend/internal/api"
	"github.com/scienceandcode/habits-todo-backend/internal/api/errors"
	"github.com/scienceandcode/habits-todo-backend/internal/api/logger"
	"github.com/scienceandcode/habits-todo-backend/internal/db"
	"github.com/scienceandcode/habits-todo-backend/internal/repository"
)

func AdminUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := errors.NewSimpleError("you are not allowed to access this resource")

		authenticatedUserId, exists := c.Get("user_id")

		if !exists || authenticatedUserId == "" {
			api.ResponseForbidden(c, err)
			logger.Error("User ID is not set in context")
			c.Abort()
			return
		}

		authenticatedUserIdAsUint, ok := authenticatedUserId.(uint)

		if !ok {
			api.ResponseForbidden(c, err)
			logger.Error("Error casting authenticatedUserId to uint")
			c.Abort()
			return
		}

		userRepository := repository.NewUserRepository(db.GetConnection())
		user, repoErr := userRepository.Repository.FindByID(authenticatedUserIdAsUint)

		if repoErr != nil {
			api.ResponseForbidden(c, err)
			logger.Error("Error finding user by ID")
			c.Abort()
			return
		}

		if !user.IsAdmin() {
			api.ResponseForbidden(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
