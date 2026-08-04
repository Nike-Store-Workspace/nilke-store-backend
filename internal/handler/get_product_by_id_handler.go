package handler

import (
	"fmt"
	"net/http"
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIDHandler struct {
	service *services.ProductService
}

func NewGetProductByIDHandler(service *services.ProductService) *GetProductByIDHandler {
	return &GetProductByIDHandler{service: service}
}

func (h *GetProductByIDHandler) GetProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	lang := ctx.DefaultQuery("lang", "fa")

	query := domain.ProductQuery{
		Lang: lang,
	}

	parsedID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {

		fmt.Println("id is not valid!")
		return
	}

	productID := uint(parsedID)

	product, err := h.service.GetById(ctx.Request.Context(), query, productID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.ErrorResponse("StatusNotFound", "Product does not exist. Please check the id and try again."))
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse(product, "The product retrieved successfully."))
}
