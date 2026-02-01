package main

import (
	"log"

	"github.com/zakoraa/golang-e-commerce-api/internal/config"
	"github.com/zakoraa/golang-e-commerce-api/internal/database"
	"github.com/zakoraa/golang-e-commerce-api/internal/router"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils"

	"go.uber.org/zap"
)

func main(){
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
	_ = db

	r := router.New()

	logger.Info("server started", zap.String("port", cfg.AppPort))
	r.Run(":" + cfg.AppPort)

}