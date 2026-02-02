package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zakoraa/golang-e-commerce-api/internal/domain/auth/usecase"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils/logger"
	"github.com/zakoraa/golang-e-commerce-api/internal/utils/response"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req usecase.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.LogError("REGISTER_USER", err)
		response.Error(
			c,
			"invalid request body",
			http.StatusBadRequest,
			"INVALID_REQUEST",
		)
		return
	}

	if err := h.authUC.Register(c.Request.Context(), req); err != nil {
		logger.LogError("REGISTER_USER", err)
		response.Error(
			c,
			err.Error(),
			http.StatusBadRequest,
			"REGISTER_FAILED",
		)
		return
	}

	logger.LogSuccess("REGISTER_USER", req.Email)
	response.SuccessNoData(
		c,
		"user registered successfully",
		http.StatusCreated,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "invalid request", http.StatusBadRequest, "INVALID_REQUEST")
		return
	}

	token, refreshToken, err := h.authUC.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, "invalid credentials", http.StatusUnauthorized, "INVALID_CREDENTIALS")
		return
	}

	c.SetCookie(
		"refresh_token",
		refreshToken,
		7*24*60*60,
		"/",
		"",
		false,
		true,
	)

	response.Success(c, token, "login success", http.StatusOK)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		response.Error(c, "unauthorized", http.StatusUnauthorized, "NO_REFRESH_TOKEN")
		return
	}

	token, err := h.authUC.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		response.Error(c, "unauthorized", http.StatusUnauthorized, "INVALID_REFRESH_TOKEN")
		return
	}

	response.Success(c, token, "token refreshed", http.StatusOK)
}
