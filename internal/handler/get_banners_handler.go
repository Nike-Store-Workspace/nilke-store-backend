package handler

import (
	"nike_store_api/internal/domain"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
)

type BannersHandler struct {
	bannersService *services.BannersService
}

func NewBannersHandler(bannersService *services.BannersService) *BannersHandler {
	return &BannersHandler{
		bannersService: bannersService,
	}
}

func (h *BannersHandler) GetBanners(c *gin.Context) {
	var query domain.BannerQuery
	lang := c.Query("lang")
	query.Lang = lang

	if query.Lang == "" {
		c.JSON(400, domain.ErrorResponse("LangRequired", "lang query parameter is required"))
		return
	}

	banners, err := h.bannersService.GetBanners(c.Request.Context(), query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var bannerResponses []domain.BannerResponse
	for _, banner := range banners {
		bannerResponse := domain.BannerResponse{
			ID:      banner.ID,
			Name:    banner.Name,
			Address: banner.Image,
		}
		bannerResponses = append(bannerResponses, bannerResponse)
	}

	c.JSON(200, domain.SuccessResponse(&bannerResponses, "Banners fetched successfully"))
}
