package services

import (
	"context"
	"nike_store_api/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type SignupService struct {
	userRepo    domain.UserRepository
	jwtSecret   string
	authService *AuthService
}

func NewSignupService(userRepo domain.UserRepository, authService *AuthService, jwtSecret string) *SignupService {
	return &SignupService{
		userRepo:    userRepo,
		jwtSecret:   jwtSecret,
		authService: authService,
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

	accessToken, err := s.authService.GenerateAccessToken(ctx, newUser)
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	refreshToken, err := s.authService.IssueRefreshToken(ctx, int(newUser.ID))
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	return &domain.AuthResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         *newUser,
	}, nil
}
