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
		logger.LogError("LOGIN_USER", err)
		response.Error(
			c,
			"invalid request body",
			http.StatusBadRequest,
			"INVALID_REQUEST",
		)
		return
	}

	token, sessionID, err := h.authUC.Login(
		c.Request.Context(),
		req,
	)
	if err != nil {
		logger.LogError("LOGIN_USER", err)
		response.Error(
			c,
			"invalid credentials",
			http.StatusUnauthorized,
			"INVALID_CREDENTIALS",
		)
		return
	}

	c.SetCookie(
		"session_id",
		sessionID,
		7*24*60*60, 
		"/",
		"",
		true,  
		true,  
	)

	logger.LogSuccess("LOGIN_USER token: ", token)
	logger.LogSuccess("LOGIN_USER sessionID: ", sessionID)
	response.Success(
		c,
		token,
		"login success",
		http.StatusOK,
	)
}