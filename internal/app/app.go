package app

import (
	"github.com/gin-gonic/gin"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/user/repository"
	"github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/usecase"
	"github.com/zakoraa/golang-e-commerce-api/internal/handler"
	"github.com/zakoraa/golang-e-commerce-api/internal/router"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"

	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	userRepo := repository.NewUserRepository(db)
	jwtManager := utils.NewJWTManagerFromEnv() 
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtManager)
	authHandler := handler.NewAuthHandler(authUsecase)

	r := gin.Default()
	router.RegisterAuth(r, authHandler)

	return &App{Router: r}
}
