package router

import (
	"github.com/zakoraa/golang-e-commerce-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.Engine, handler *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
