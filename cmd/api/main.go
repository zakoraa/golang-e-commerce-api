package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/zakoraa/golang-e-commerce-api/internal/config"
)

func main (){
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("service running on port", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}