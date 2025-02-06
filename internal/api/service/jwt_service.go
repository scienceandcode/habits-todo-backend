package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

type JWTService struct{}

func (j *JWTService) GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(common.GetEnv("JWT_SECRET_KEY")))
}

func (j *JWTService) ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(common.GetEnv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user_id in token")
		}
		return uint(userID), nil
	}

	return 0, errors.New("invalid token")
}

func NewJWTService() *JWTService {
	return &JWTService{}
}
