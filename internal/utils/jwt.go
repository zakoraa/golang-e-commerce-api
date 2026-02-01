package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret string
}

func NewJWTManagerFromEnv() *JWTManager {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return &JWTManager{secret}
}

func (j *JWTManager) Generate(userID, role string) (string, string) {
	accessClaims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).
		SignedString([]byte(j.secret))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(j.secret))

	return accessToken, refreshToken
}
