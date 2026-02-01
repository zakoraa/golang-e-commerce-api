package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine{
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}

func RegisterAuth(
	r *gin.Engine,
	handler *handler.AuthHandler,
) {
	auth := r.Group("/auth"){
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}