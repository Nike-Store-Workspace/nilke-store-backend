package handler

import (
	"errors"
	"fmt"
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type SignupHandler struct {
	service services.SignupService
}

func NewSignupHandler(service *services.SignupService) *SignupHandler {
	return &SignupHandler{service: *service}
}

func (h *SignupHandler) Signup(c *gin.Context) {
	var request domain.SignupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse(
			"StatusBadRequest",
			fmt.Sprintf("Invalid request body.\n details: %s", err.Error()),
		))
		return
	}

	resp, err := h.service.Signup(c.Request.Context(), request)

	if err != nil {
		// ۲. بررسی نوع خطاهای بزنسی و اختصاص استاتوس کد مناسب
		switch {
		case errors.Is(err, domain.ErrPasswordMismatch):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400
		case errors.Is(err, domain.ErrUserAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"}) // 500
		}
		return
	}

	c.JSON(http.StatusCreated, resp)

}
