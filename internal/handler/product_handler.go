package handler

import (
	"fmt"
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	lang := c.DefaultQuery("lang", "fa")
	sort := c.DefaultQuery("sort", "newest")
	category := c.Query("category")

	query := domain.ProductQuery{
		Lang:     lang,
		Sort:     sort,
		Category: category,
	}

	products, err := h.service.GetProducts(c.Request.Context(), query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("An Internal server accorded in get products : %w", err.Error())})
	}

	c.JSON(
		http.StatusOK, gin.H{
			"data": products,
		})
}
