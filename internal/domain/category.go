package domain

import "context"

type Category struct {
	ID     int64  `json:"id"`
	Slug   string `json:"slug"`
	NameEn string `json:"name_en"`
	NameFa string `json:"name_fa"`
}

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]Category, error)
}
