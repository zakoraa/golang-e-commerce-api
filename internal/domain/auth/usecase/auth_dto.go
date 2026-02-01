package usecase

import "github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"

type RegisterRequest struct {
	Email string
	Password string
	Role entity.UserRole
}

type LoginRequest struct {
	Email string
	Password string
}

type TokenResponse struct {
	AccessToken string
	RefreshToken string
}
