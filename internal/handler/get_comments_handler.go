package handler

import (
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCommentsHandler struct {
	service *services.CommentService
}

func NewGetCommentHandler(service services.CommentService) *GetCommentsHandler {
	return &GetCommentsHandler{service: &service}
}

func (h *GetCommentsHandler) GetProductComments(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	lang := c.DefaultQuery("lang", "fa")

	query := domain.CommentQuery{
		ProductID: productId,
		Page:      page,
		Limit:     limit,
		Lang:      lang,
	}

	comments, err := h.service.GetByProductID(c.Request.Context(), query)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"limit": query.Limit,
		"lang":  query.Lang,
		"data":  comments,
	})
}

func (h *GetCommentsHandler) AddComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
	}

	userID := userIDVal.(int64)

	var req domain.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your requested data is invalid", "details": err.Error()})
		return
	}

	comment := domain.Comment{
		ProductID: req.ProductID,
		UserID:    userID,
		TitleEn:   req.TitleEn,
		TitleFa:   req.TitleFa,
		Body:      req.Body,
		Rating:    req.Rating,
	}

	if err := h.service.Create(c.Request.Context(), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "comment saved successfully",
		"data":    comment,
	})
}
