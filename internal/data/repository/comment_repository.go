package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"nike_store_api/internal/domain"
)

type PostgresCommentRepository struct {
	db *sql.DB
}

func NewPostgresCommentRepository(db *sql.DB) domain.CommentRepository {
	return &PostgresCommentRepository{db: db}
}

func (r *PostgresCommentRepository) GetByProductID(ctx context.Context, query domain.CommentQuery) ([]domain.Comment, error) {
	bodyColumn := "c.body_fa"
	if query.Lang == "en" {
		bodyColumn = "c.body_en"
	}

	sqlQuery := fmt.Sprintf(`
		SELECT c.id, c.product_id, c.user_id, u.full_name,c.title_en, c.title_fa, %s, c.rating, c.created_at
		FROM product_comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.product_id = $1
		ORDER BY c.created_at DESC
		LIMIT $2 OFFSET $3
	`, bodyColumn)

	rows, err := r.db.QueryContext(ctx, sqlQuery, query.ProductID, query.Limit, query.Offset)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %w", err)
	}
	defer rows.Close()

	comments := []domain.Comment{}
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(
			&c.ID,
			&c.ProductID,
			&c.UserID,
			&c.UserName,
			&c.TitleEn,
			&c.TitleFa,
			&c.Body,
			&c.Rating,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *PostgresCommentRepository) Create(ctx context.Context, comment *domain.Comment) error {

	var exists bool
	checkProductQuery := `SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)`
	err := r.db.QueryRowContext(ctx, checkProductQuery, comment.ProductID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking product existence: %w", err)
	}

	if !exists {
		return errors.New("Product does not exist!")
	}

	query := `
		INSERT INTO product_comments (product_id, user_id, title_en,title_fa, body_fa, body_en, rating)
		VALUES ($1, $2, $3, $4, $5, $6,$7)
		RETURNING id, created_at
	`
	// متن کامنت کاربر را هم برای body_fa و هم body_en ذخیره می‌کنیم
	return r.db.QueryRowContext(
		ctx, query,
		comment.ProductID, comment.UserID, comment.TitleEn, comment.TitleFa, comment.Body, comment.Body, comment.Rating,
	).Scan(&comment.ID, &comment.CreatedAt)
}
