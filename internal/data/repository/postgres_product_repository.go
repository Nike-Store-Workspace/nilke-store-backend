package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
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

	err, products := r.scanProducts(ctx, rows)

	return products, nil

}

func (r *PostgresProductRepository) GetById(ctx context.Context, query domain.ProductQuery, id uint) (domain.Product, error) {

	var p domain.Product

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
		&p.CategorySlug,
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
func (r *PostgresProductRepository) Search(
	ctx context.Context,
	query domain.ProductQuery,
	searchTerm string,
) ([]domain.Product, error) {

	var dbQuery string

	if query.Lang == "fa" {
		dbQuery = `
			SELECT 
			    p.id,
			    p.category_id,
			    p.title_en,
			    p.title_fa,
			    p.description_en,
			    p.description_fa,
			    p.price_toman,
			    p.price_usd,
			    p.previous_price_toman,
			    p.previous_price_usd,
			    p.discount_percentage,
			    p.images,
			    p.created_at,
			    c.slug
			FROM products p 
			LEFT JOIN categories c ON p.category_id = c.id  
			WHERE p.title_fa ILIKE '%' || $1 || '%'
		`
	} else {
		dbQuery = `
			SELECT 
			    p.id,
			    p.category_id,
			    p.title_en,
			    p.title_fa,
			    p.description_en,
			    p.description_fa,
			    p.price_toman,
			    p.price_usd,
			    p.previous_price_toman,
			    p.previous_price_usd,
			    p.discount_percentage,
			    p.images,
			    p.created_at,
			    c.slug
			FROM products p 
			LEFT JOIN categories c ON p.category_id = c.id  
			WHERE p.title_en ILIKE '%' || $1 || '%'
		`
	}

	rows, err := r.db.QueryContext(ctx, dbQuery, searchTerm)
	if err != nil {
		return nil, errors.New("error in search process: " + err.Error())
	}
	defer rows.Close()

	err, products := r.scanProducts(ctx, rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *PostgresProductRepository) scanProducts(ctx context.Context, rows *sql.Rows) (error, []domain.Product) {

	var products []domain.Product
	var productIDs []uint

	for rows.Next() {
		var p domain.Product

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
			&p.CategorySlug,
		)
		// ...

		if err != nil {
			return errors.New("error in process " + err.Error()), nil
		}

		products = append(products, p)
		productIDs = append(productIDs, p.ID)

	}

	if len(products) == 0 {
		return nil, []domain.Product{}
	}

	variantSQL := `SELECT id, product_id, color_en, color_fa, size, stock FROM product_variants WHERE product_id = ANY($1)`

	var pIDsInt64 []int64
	for _, id := range productIDs {
		pIDsInt64 = append(pIDsInt64, int64(id))
	}

	variantRows, err := r.db.QueryContext(ctx, variantSQL, pq.Array(pIDsInt64))
	if err != nil {
		return fmt.Errorf("error querying variants: %w", err), nil
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
			return fmt.Errorf("error scanning variant: %w", err), nil
		}

		variantMap[variant.ProductID] = append(variantMap[variant.ProductID], variant)
	}

	for i := range products {
		products[i].Variants = variantMap[products[i].ID]
	}

	return nil, products
}
