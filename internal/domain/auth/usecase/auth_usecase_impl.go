package usecase

import (
	"time"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	userEntity "github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"
	authEntity "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/entity"
	userRepo "github.com/zakoraa/golang-e-commerce-api/internal/domain/user/repository"
	authRepo "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/repository"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"
)

type authUsecase struct {
	userRepo   userRepo.UserRepository
	refreshRepo authRepo.RefreshTokenRepository
	jwt *utils.JWTManager
}

func NewAuthUsecase(
	userRepo userRepo.UserRepository,
	refreshRepo authRepo.RefreshTokenRepository,
	jwt *utils.JWTManager,
) AuthUsecase {
	return &authUsecase{userRepo, refreshRepo, jwt}
}

func (u *authUsecase) Register (ctx context.Context, req RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &userEntity.User {
		ID: uuid.New(),
		Email: req.Email,
		Password: string(hash),
		Role: req.Role,
	}

	return u.userRepo.Create(ctx, user)
}

func (u *authUsecase) Login(
	ctx context.Context,
	req LoginRequest,
) (TokenResponse, string, error) {

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return TokenResponse{}, "", errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	) != nil {
		return TokenResponse{}, "", errors.New("invalid credentials")
	}

	accessToken, err := u.jwt.GenerateAccessToken(
		user.ID.String(),
		string(user.Role),
	)
	if err != nil {
		return TokenResponse{}, "", err
	}

	refreshToken := &authEntity.RefreshToken{
		ID: uuid.New().String(),
		UserID: user.ID.String(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := u.refreshRepo.Create(ctx, refreshToken); err != nil {
		return TokenResponse{}, "", err
	}

	return TokenResponse{
		AccessToken: accessToken,
	}, refreshToken.ID, nil
}

func (u *authUsecase) Refresh(
	ctx context.Context,
	refreshTokenID string,
) (TokenResponse, error) {

	token, err := u.refreshRepo.FindValid(ctx, refreshTokenID, time.Now())
	if err != nil {
		return TokenResponse{}, errors.New("invalid refresh token")
	}

	user, err := u.userRepo.FindByID(ctx, token.UserID)
	if err != nil {
		return TokenResponse{}, err
	}

	accessToken, err := u.jwt.GenerateAccessToken(
		user.ID.String(),
		string(user.Role),
	)
	if err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		AccessToken: accessToken,
	}, nil
}
