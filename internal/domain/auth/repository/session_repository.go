package repository

import (
	"context"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/entity"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
}
