package handler

import (
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authService.Login(c.Request.Context(), request)

	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			})
		return

	}
	c.JSON(http.StatusOK, domain.SuccessResponse(resp, "You are login successfully."))
}
