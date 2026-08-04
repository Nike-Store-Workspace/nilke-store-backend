package handler

import (
	"fmt"
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type SearchProductsHandler struct {
	service *services.ProductService
}

func NewSearchProductsHandler(service *services.ProductService) *SearchProductsHandler {
	return &SearchProductsHandler{service: service}
}

func (h *SearchProductsHandler) Search(c *gin.Context) {
	lang := c.DefaultQuery("lang", "fa")
	sort := c.DefaultQuery("sort", "newest")
	category := c.Query("category")
	searchTerm := c.Query("s")

	query := domain.ProductQuery{
		Lang:     lang,
		Sort:     sort,
		Category: category,
	}

	products, err := h.service.Search(c.Request.Context(), query, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse("INTERNAL_ERROR", fmt.Sprintf("An Internal server error accorded in get products : %w", err.Error())))
		return
	}

	c.JSON(
		http.StatusOK, domain.SuccessResponse(&products, "Products retrieved successfully"))
}
