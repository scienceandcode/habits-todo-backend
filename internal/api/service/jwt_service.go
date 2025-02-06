package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	SecretKey string
}

func (j *JWTService) GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.FormatUint(uint64(userID), 10),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWTService) ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"]
		if !ok {
			return 0, errors.New("invalid user_id in token")
		}

		parsedUserID, _ := strconv.ParseUint(userID.(string), 10, 64)
		return uint(parsedUserID), nil
	}

	return 0, errors.New("invalid token")
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
	}
}
