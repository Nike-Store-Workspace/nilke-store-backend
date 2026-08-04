package services

import (
	"context"
	"nike_store_api/internal/domain"
)

type ProductService struct {
	productRepository domain.ProductRepository
}

func NewProductService(productRepository domain.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) GetProducts(ctx context.Context, query domain.ProductQuery) ([]domain.ProductResponse, error) {
	products, err := s.productRepository.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}

	return createProductsResponse(products, query)
}

func (s *ProductService) GetById(ctx context.Context, query domain.ProductQuery, id uint) (*domain.ProductResponse, error) {
	product, err := s.productRepository.GetById(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var title string
	var desc string

	if query.Lang == "fa" {
		title = product.TitleFa
		desc = product.DescriptionFa
	} else {
		title = product.TitleEn
		desc = product.DescriptionEn
	}

	var variantResponses []domain.ProductVariantResponse
	for _, v := range product.Variants {
		color := v.ColorEn
		if query.Lang == "fa" {
			color = v.ColorFa
		}
		variantResponses = append(variantResponses, domain.ProductVariantResponse{
			ID:    v.ID,
			Color: color,
			Size:  v.Size,
			Stock: v.Stock,
		})
	}

	response := domain.ProductResponse{
		ID:                 product.ID,
		Category:           product.CategorySlug,
		Title:              title,
		Description:        desc,
		PriceUSD:           product.PriceUSD,
		PriceToman:         product.PriceToman,
		PreviousPriceUSD:   product.PreviousPriceUSD,
		PreviousPriceToman: product.PreviousPriceToman,
		DiscountPercentage: product.DiscountPercentage,
		Images:             product.Images,
		Variants:           variantResponses,
	}
	return &response, nil

}
func (s *ProductService) Search(ctx context.Context, query domain.ProductQuery, searchTerm string) ([]domain.ProductResponse, error) {
	products, err := s.productRepository.Search(ctx, query, searchTerm)
	if err != nil {
		return nil, err
	}

	return createProductsResponse(products, query)

}

func createProductsResponse(products []domain.Product, query domain.ProductQuery) ([]domain.ProductResponse, error) {
	if len(products) == 0 {
		return []domain.ProductResponse{}, nil
	}

	var response []domain.ProductResponse

	for _, product := range products {
		var title string
		var desc string

		if query.Lang == "fa" {
			title = product.TitleFa
			desc = product.DescriptionFa
		} else {
			title = product.TitleEn
			desc = product.DescriptionEn
		}

		var variantResponses []domain.ProductVariantResponse
		for _, v := range product.Variants {
			color := v.ColorEn
			if query.Lang == "fa" {
				color = v.ColorFa
			}
			variantResponses = append(variantResponses, domain.ProductVariantResponse{
				ID:    v.ID,
				Color: color,
				Size:  v.Size,
				Stock: v.Stock,
			})
		}

		mappedProduct := domain.ProductResponse{
			ID:                 product.ID,
			Category:           product.CategorySlug,
			Title:              title,
			Description:        desc,
			PriceUSD:           product.PriceUSD,
			PriceToman:         product.PriceToman,
			PreviousPriceUSD:   product.PreviousPriceUSD,
			PreviousPriceToman: product.PreviousPriceToman,
			DiscountPercentage: product.DiscountPercentage,
			Images:             product.Images,
			Variants:           variantResponses,
		}

		response = append(response, mappedProduct)
	}

	return response, nil
}
