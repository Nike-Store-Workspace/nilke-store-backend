package handler

import (
	"fmt"
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type GetAllProductHandler struct {
	service *services.ProductService
}

func NewGetAllProductHandler(service *services.ProductService) *GetAllProductHandler {
	return &GetAllProductHandler{service: service}
}

func (h *GetAllProductHandler) GetProducts(c *gin.Context) {

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
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse("INTERNAL_ERROR", fmt.Sprintf("An Internal server error accorded in get products : %w", err.Error())))
		return
	}

	c.JSON(
		http.StatusOK, domain.SuccessResponse(&products, "Products retrieved successfully"))

}
