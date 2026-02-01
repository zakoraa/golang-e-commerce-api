package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/entity"
	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/repository"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"
)

type authUsecase struct {
	userRepo repository.UserRepository
	jwt *utils.JWTManager
}

func NewAuthUsecase(userRepo repository.UserRepository, jwt *utils.JWTManager) AuthUsecase {
	return &authUsecase{userRepo, jwt}
}

func (u *authUsecase) Register (ctx context.Context, req RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	user := &entity.User {
		ID: uuid.New(),
		Email: req.Email,
		Password: string(hash),
		Role: req.Role,
	}

	return u.userRepo.Create(ctx, user)
}

func (u *authUsecase) Login (ctx context.Context, req LoginRequest) (TokenResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return TokenResponse{}, errors.New("Invalid credentials")
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	) != nil {
		return TokenResponse{}, errors.New("Invalid credentials")
	}

	access, refresh := u.jwt.Generate(user.ID.String(), string(user.Role))

	return TokenResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}, nil
}
