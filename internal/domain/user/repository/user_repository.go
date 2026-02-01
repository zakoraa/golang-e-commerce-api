package repository

import (
	"context"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindById(ctx context.Context, id string) (*entity.User, error)
}