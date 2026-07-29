package services

import (
	"context"
	"nike_store_api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SignupService struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewSignupService(userRepo domain.UserRepository, jwtSecret string) *SignupService {
	return &SignupService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *SignupService) Signup(ctx context.Context, request domain.SignupRequest) (*domain.AuthResponse, error) {
	existingUser, _ := s.userRepo.GetByEmail(ctx, string(request.Email))
	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists // <-- 409 Conflict
	}

	if request.Password != request.ConfirmPass {
		return nil, domain.ErrPasswordMismatch // <-- 400 Bad Request
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.ErrInternalServer // <-- 500
	}

	newUser := &domain.User{
		Email:        request.Email,
		PasswordHash: string(hashPassword),
		FullName:     request.FullName,
	}

	err = s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, domain.ErrInternalServer // <-- 500
	}

	claims := jwt.MapClaims{
		"user_id": newUser.ID,
		"email":   newUser.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ⚠️ استفاده از HS256 برای Secret متنی
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, domain.ErrInternalServer // <-- 500
	}

	return &domain.AuthResponse{
		Token: tokenString,
		User:  *newUser,
	}, nil
}
