package repository

import (
	"context"
	"database/sql"
	"nike_store_api/internal/domain"
)

type BannerRepository struct {
	db *sql.DB
}

func NewBannerRepository(db *sql.DB) *BannerRepository {
	return &BannerRepository{
		db: db,
	}
}

func (r *BannerRepository) GetBanners(ctx context.Context, lang domain.BannerQuery) ([]domain.Banner, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, image, lang FROM banners WHERE lang = $1", lang.Lang)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []domain.Banner
	for rows.Next() {
		var banner domain.Banner
		err := rows.Scan(&banner.ID, &banner.Name, &banner.Image, &banner.Lang)
		if err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}

	return banners, nil
}
