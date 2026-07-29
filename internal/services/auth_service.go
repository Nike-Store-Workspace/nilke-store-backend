package services

import (
	"context"
	"errors"
	"fmt"
	"nike_store_api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewAuthService(
	userRepo domain.UserRepository,
	jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, request domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, request.Email)

	if err != nil {
		return nil, errors.New("email or password is incorrect!")
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	fmt.Println(string(pass))

	fmt.Printf("۱. پسورد دریافتی از JSON: [%s]\n", request.Password)
	fmt.Printf("۲. هش دریافتی از دیتابیس: [%s] (طول: %d)\n", user.PasswordHash, len(user.PasswordHash))
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return nil, errors.New("email or password is incorrect!")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))

	if err != nil {

		return nil, errors.New("error in generating token!")

	}

	return &domain.AuthResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}
