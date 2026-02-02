package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	authEntity "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/entity"
	authRepo "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/repository"
)

type sessionRepository struct {
	rdb *redis.Client
}

func NewSessionRepository(rdb *redis.Client) authRepo.SessionRepository {
	return &sessionRepository{rdb: rdb}
}

func (r *sessionRepository) Create(
	ctx context.Context,
	session *authEntity.Session,
) error {
	return r.rdb.Set(
		ctx,
		session.ID,
		session.UserID,
		time.Until(session.ExpiresAt),
	).Err()
}
