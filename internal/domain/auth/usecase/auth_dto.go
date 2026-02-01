package usecase

import "github.com/zakoraa/e-commerce-api/internal/domain/user/entity"

type RegisterRequest struct {
	Email    string
	Password string
	Role     UserRole
}

type LoginRequest struct {
	Email    string
	Password string
}

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
}
