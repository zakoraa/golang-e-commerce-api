package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repo "github.com/zakoraa/golang-e-commerce-api/internal/database"
	authUsecase "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/usecase"
	"github.com/zakoraa/golang-e-commerce-api/internal/handler"
	"github.com/zakoraa/golang-e-commerce-api/internal/router"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"
)

type App struct {
	Router *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	userRepository := repo.NewUserRepository(db)
	refreshTokenRepository := repo.NewRefreshTokenRepository(db)

	jwtManager := utils.NewJWTManagerFromEnv()

	authUC := authUsecase.NewAuthUsecase(
		userRepository,
		refreshTokenRepository,
		jwtManager,
	)

	authHandler := handler.NewAuthHandler(authUC)

	r := gin.Default()
	router.RegisterAuth(r, authHandler)

	return &App{Router: r}
}
