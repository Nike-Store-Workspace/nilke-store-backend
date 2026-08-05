package domain

import "context"

type Banner struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Lang  string `json:"lang"`
}

type BannerResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BannerQuery struct {
	Lang string
}

type BannerRepository interface {
	GetBanners(ctx context.Context, query BannerQuery) ([]Banner, error)
}
