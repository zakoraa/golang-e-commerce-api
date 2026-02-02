package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	userRepo "github.com/zakoraa/golang-e-commerce-api/internal/database"
	sessionRepo "github.com/zakoraa/golang-e-commerce-api/internal/cache"
	authUsecase "github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/usecase"

	"github.com/zakoraa/golang-e-commerce-api/internal/handler"
	"github.com/zakoraa/golang-e-commerce-api/internal/router"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"
)

type App struct {
	Router *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	userRepository := userRepo.NewUserRepository(db)
	redisClient := utils.NewRedisClient()
	sessionRepository := sessionRepo.NewSessionRepository(redisClient)

	jwtManager := utils.NewJWTManagerFromEnv()

	authUC := authUsecase.NewAuthUsecase(
		userRepository,
		sessionRepository,
		jwtManager,
	)

	authHandler := handler.NewAuthHandler(authUC)

	r := gin.Default()
	router.RegisterAuth(r, authHandler)

	return &App{Router: r}
}
