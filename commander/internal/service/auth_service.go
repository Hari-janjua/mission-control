package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

var jwtKey = []byte("supersecretkey")

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) GenerateToken(soldierID string) (string, error) {
	claims := &jwt.MapClaims{
		"soldier_id": soldierID,
		"exp":        time.Now().Add(30 * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (a *AuthService) ValidateToken(tokenStr string) (bool, string) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return false, ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, claims["soldier_id"].(string)
	}
	return false, ""
}
