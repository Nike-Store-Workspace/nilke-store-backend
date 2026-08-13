package repository

import (
	"context"
	"database/sql"
	"errors"
	"nike_store_api/internal/domain"
)

type PostgresRefreshTokenRepository struct {
	db *sql.DB
}

func NewPostgresRefreshTokenRepository(db *sql.DB) domain.RefreshTokenRepository {
	return &PostgresRefreshTokenRepository{db: db}
}

func (r *PostgresRefreshTokenRepository) CreateRefreshToken(ctx context.Context, refreshToken *domain.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id,token_hash,expires_at) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, refreshToken.UserID, refreshToken.TokenHash, refreshToken.ExpiresAt).Scan(&refreshToken.ID, &refreshToken.CreatedAt)
}

func (r *PostgresRefreshTokenRepository) GetRefreshTokenByHash(ctx context.Context, tokenHash string) (*domain.RefreshToken, error) {
	query := `SELECT id, user_id, token_hash,expires_at,revoked_at,created_at FROM refresh_tokens WHERE token_hash = $1`

	var rt domain.RefreshToken
	err := r.db.QueryRowContext(ctx, query, tokenHash).Scan(
		&rt.ID,
		&rt.UserID,
		&rt.TokenHash,
		&rt.ExpiresAt,
		&rt.RevokedAt,
		&rt.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrInvalidRefreshToken
		}
		return nil, err
	}
	return &rt, nil
}
func (r *PostgresRefreshTokenRepository) RevokeRefreshToken(ctx context.Context, id string) error {
	query := `UPDATE refresh_tokens SET revoked_at = now() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func (r *PostgresRefreshTokenRepository) RevokeAllRefreshTokensForUser(ctx context.Context, userID int64) error {
	query := `UPDATE refresh_tokens SET revoked_at = now() WHERE user_id = $1 AND revoked_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
