package main

import (
	"log"

	"github.com/zakoraa/golang-e-commerce-api/internal/config"
	"github.com/zakoraa/golang-e-commerce-api/internal/database"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"
	"github.com/zakoraa/golang-e-commerce-api/internal/app" 

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger := utils.NewLoger()
	defer logger.Sync()

	dsn := database.BuildDSN(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	db := database.NewPostgres(dsn)

	app := app.NewApp(db)

	logger.Info("server started", zap.String("port", cfg.AppPort))
	app.Router.Run(":" + cfg.AppPort)
}
