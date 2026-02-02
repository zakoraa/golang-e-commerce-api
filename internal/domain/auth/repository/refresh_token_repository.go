package repository

import (
	"context"
	"time"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/entity"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *entity.RefreshToken) error
	FindValid(ctx context.Context, tokenID string, now time.Time) (*entity.RefreshToken, error)
	Delete(ctx context.Context, tokenID string) error
	DeleteByUser(ctx context.Context, userID string) error
}
