package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	
	return &user, err
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*entity.User, error){
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error

	return &user, err
}