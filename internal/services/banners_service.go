package services

import (
	"context"
	"nike_store_api/internal/data/repository"
	"nike_store_api/internal/domain"
)

type BannersService struct {
	bannerRepo *repository.BannerRepository
}

func NewBannersService(bannerRepo *repository.BannerRepository) *BannersService {
	return &BannersService{
		bannerRepo: bannerRepo,
	}
}

func (s *BannersService) GetBanners(ctx context.Context, query domain.BannerQuery) ([]domain.Banner, error) {
	return s.bannerRepo.GetBanners(ctx, query)
}
