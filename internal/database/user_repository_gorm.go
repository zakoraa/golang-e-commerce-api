package database

import (
	"context"

	"gorm.io/gorm"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"
	userRepo "github.com/zakoraa/golang-e-commerce-api/internal/domain/user/repository"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepo.UserRepository {
	return &userRepoImpl{db: db}
}

func (r *userRepoImpl) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepoImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	return &user, err
}

func (r *userRepoImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error

	return &user, err
}
