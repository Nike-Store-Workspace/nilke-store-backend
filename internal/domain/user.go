package domain

import "context"

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	FullName     string `json:"full_name"`
	CreatedAt    string `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type SignupRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	FullName    string `json:"full_name" binding:"required,min=2,max=100"`
	ConfirmPass string `json:"confirm_password" binding:"required,min=6"`
}
type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
