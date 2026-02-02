package database

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/entity"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *refreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) Create(
	ctx context.Context,
	token *entity.RefreshToken,
) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *refreshTokenRepository) FindValid(
	ctx context.Context,
	tokenID string,
	now time.Time,
) (*entity.RefreshToken, error) {
	var token entity.RefreshToken
	err := r.db.WithContext(ctx).
		Where("id = ? AND expires_at > ?", tokenID, now).
		First(&token).Error

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *refreshTokenRepository) Delete(
	ctx context.Context,
	tokenID string,
) error {
	return r.db.WithContext(ctx).
		Delete(&entity.RefreshToken{}, "id = ?", tokenID).
		Error
}

func (r *refreshTokenRepository) DeleteByUser(
	ctx context.Context,
	userID string,
) error {
	return r.db.WithContext(ctx).
		Delete(&entity.RefreshToken{}, "user_id = ?", userID).
		Error
}
