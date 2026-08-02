package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"nike_store_api/internal/domain"

	"github.com/lib/pq"
)

//go:embed queries/get_products.sql
var baseProductSql string

//go:embed queries/get_product_by_id.sql
var getProductByIdQuery string

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) domain.ProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) GetAll(ctx context.Context, query domain.ProductQuery) ([]domain.Product, error) {
	baseQuery := baseProductSql

	var args []interface{}
	argCounter := 1

	if query.Category != "" {
		baseQuery += fmt.Sprintf(" AND c.slug = $%d", argCounter)
		args = append(args, query.Category)
		argCounter++
	}
	switch query.Sort {
	case "price_asc":
		baseQuery += " ORDER BY p.price_toman ASC"
	case "price_desc":
		baseQuery += " ORDER BY p.price_toman DESC"
	case "newest":
		baseQuery += " ORDER BY p.created_at DESC"
	default:
		baseQuery += " ORDER BY p.created_at DESC"
	}

	// run query
	rows, err := r.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying products: %w", err)
	}

	defer rows.Close()

	var products []domain.Product
	var productIDs []uint

	for rows.Next() {
		var p domain.Product
		var categorySlug sql.NullString

		// ...
		err := rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.TitleEn,
			&p.TitleFa,
			&p.DescriptionEn,
			&p.DescriptionFa,
			&p.PriceToman, // <--- اول تومان (مطابق فایل SQL)
			&p.PriceUSD,   // <--- بعد دلار (مطابق فایل SQL)
			&p.PreviousPriceToman,
			&p.PreviousPriceUSD,
			&p.DiscountPercentage,
			pq.Array(&p.Images),
			&p.CreatedAt,
			&categorySlug,
		)
		// ...

		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		products = append(products, p)
		productIDs = append(productIDs, p.ID)

	}

	if len(products) == 0 {
		return []domain.Product{}, nil
	}

	variantSQL := `SELECT id, product_id, color_en, color_fa, size, stock FROM product_variants WHERE product_id = ANY($1)`

	var pIDsInt64 []int64
	for _, id := range productIDs {
		pIDsInt64 = append(pIDsInt64, int64(id))
	}

	variantRows, err := r.db.QueryContext(ctx, variantSQL, pq.Array(pIDsInt64))
	if err != nil {
		return nil, fmt.Errorf("error querying variants: %w", err)
	}

	defer variantRows.Close()

	variantMap := make(map[uint][]domain.ProductVariant)

	for variantRows.Next() {
		var variant domain.ProductVariant
		err := variantRows.Scan(
			&variant.ID,
			&variant.ProductID,
			&variant.ColorEn,
			&variant.ColorFa,
			&variant.Size,
			&variant.Stock,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning variant: %w", err)
		}

		variantMap[variant.ProductID] = append(variantMap[variant.ProductID], variant)
	}

	for i := range products {
		products[i].Variants = variantMap[products[i].ID]
	}

	return products, nil

}

func (r *PostgresProductRepository) GetById(ctx context.Context, query domain.ProductQuery, id uint) (domain.Product, error) {

	var p domain.Product
	var categorySlug sql.NullString

	row := r.db.QueryRowContext(ctx, getProductByIdQuery, id)

	err := row.Scan(
		&p.ID,
		&p.CategoryID,
		&p.TitleEn,
		&p.TitleFa,
		&p.DescriptionEn,
		&p.DescriptionFa,
		&p.PriceToman,
		&p.PriceUSD,
		&p.PreviousPriceToman,
		&p.PreviousPriceUSD,
		&p.DiscountPercentage,
		pq.Array(&p.Images),
		&p.CreatedAt,
		&categorySlug,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, fmt.Errorf("product not found")
		}
		return domain.Product{}, fmt.Errorf("error scanning product by id: %w", err)
	}

	variantSQL := `SELECT id, product_id, color_en, color_fa, size, stock FROM product_variants WHERE product_id = $1`
	variantRows, err := r.db.QueryContext(ctx, variantSQL, p.ID)

	if err != nil {
		return domain.Product{}, fmt.Errorf("error querying variants for product: %w", err)
	}

	defer variantRows.Close()

	var variants []domain.ProductVariant
	for variantRows.Next() {

		var v domain.ProductVariant
		err := variantRows.Scan(&v.ID, &v.ProductID, &v.ColorEn, &v.ColorFa, &v.Size, &v.Stock)
		if err != nil {
			return domain.Product{}, fmt.Errorf("error scanning variant: %w", err)
		}
		variants = append(variants, v)

	}
	p.Variants = variants

	return p, nil
}
