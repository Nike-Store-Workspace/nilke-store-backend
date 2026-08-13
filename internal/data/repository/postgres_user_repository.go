package repository

import (
	"context"
	"database/sql"
	"errors"
	"nike_store_api/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id,email,password_hash,full_name,created_at FROM users WHERE email = $1`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No user founded with this email!")
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (email, password_hash, full_name) VALUES ($1, $2, $3) RETURNING id, created_at;`

	return r.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash, user.FullName).Scan(&user.ID, &user.CreatedAt)
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id,email,password_hash,full_name,created_at FROM users WHERE id = $1`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No user founded with this email!")
		}
		return nil, err
	}
	return &user, nil
}
