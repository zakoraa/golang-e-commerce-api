package usecase

import "context"

type AuthUsecase interface {
	Register(ctx context.Context, req RegisterRequest) error
	Login(ctx context.Context, req LoginRequest) (TokenResponse, error)
}
