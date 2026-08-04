package domain

import (
	"context"
	"time"
)

type Product struct {
	ID                 uint             `json:"id"`
	CategoryID         uint             `json:"category_id"`
	CategorySlug       string           `json:"category_slug"`
	TitleEn            string           `json:"title_en"`
	TitleFa            string           `json:"title_fa"`
	DescriptionEn      string           `json:"description_en"`
	DescriptionFa      string           `json:"description_fa"`
	PriceUSD           float64          `json:"price_usd"`
	PriceToman         int64            `json:"price_toman"`
	PreviousPriceUSD   *float64         `json:"previous_price_usd"`
	PreviousPriceToman *int64           `json:"previous_price_toman"`
	DiscountPercentage *float64         `json:"discount_percentage"`
	Images             []string         `json:"images"`
	CreatedAt          time.Time        `json:"created_at"`
	Variants           []ProductVariant `json:"variants"`
}

type ProductResponse struct {
	ID                 uint                     `json:"id"`
	Category           string                   `json:"category"` // اینجا رشته می‌فرستیم، نه آبجکت کامل را
	Title              string                   `json:"title"`
	Description        string                   `json:"description"`
	PriceUSD           float64                  `json:"price_usd"`
	PriceToman         int64                    `json:"price_toman"`
	PreviousPriceUSD   *float64                 `json:"previous_price_usd"`
	PreviousPriceToman *int64                   `json:"previous_price_toman"`
	DiscountPercentage *float64                 `json:"discount_percentage"`
	Images             []string                 `json:"images"`
	Variants           []ProductVariantResponse `json:"variants"`
}

type ProductQuery struct {
	Lang     string
	Sort     string
	Category string
}

type ProductVariant struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ColorEn   string `json:"color_en"`
	ColorFa   string `json:"color_fa"`
	Size      string `json:"size"`
	Stock     int    `json:"stock"`
}

type ProductVariantResponse struct {
	ID    uint   `json:"id"`
	Color string `json:"color"`
	Size  string `json:"size"`
	Stock int    `json:"stock"`
}

type ProductRepository interface {
	GetAll(ctx context.Context, query ProductQuery) ([]Product, error)
	GetById(ctx context.Context, query ProductQuery, id uint) (Product, error)
	Search(ctx context.Context, query ProductQuery, searchTerm string) ([]Product, error)
}
