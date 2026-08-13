package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"nike_store_api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo         domain.UserRepository
	refreshTokenRepo domain.RefreshTokenRepository
	jwtSecret        string
}

func NewAuthService(
	userRepo domain.UserRepository,
	refreshTokenRepo domain.RefreshTokenRepository,
	jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        jwtSecret,
	}
}

const (
	accessTokenTTL  = 1 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
)

func generateSecureToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

func (s *AuthService) GenerateAccessToken(ctx context.Context, user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(accessTokenTTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))

}

func (s *AuthService) IssueRefreshToken(ctx context.Context, userId int) (string, error) {

	rawToken, err := generateSecureToken()
	if err != nil {
		return "", errors.New("error generating refresh token")

	}

	rt := &domain.RefreshToken{
		UserID:    userId,
		TokenHash: hashToken(rawToken),
		ExpiresAt: time.Now().Add(refreshTokenTTL),
	}

	if err := s.refreshTokenRepo.CreateRefreshToken(ctx, rt); err != nil {
		return "", errors.New("error saving refresh token")
	}

	return rawToken, nil

}

func (s *AuthService) Login(ctx context.Context, request domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, request.Email)

	if err != nil {
		return nil, errors.New("email or password is incorrect!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return nil, errors.New("email or password is incorrect!")
	}

	accessToken, err := s.GenerateAccessToken(ctx, user)
	if err != nil {
		return nil, errors.New("error in generating access token!")
	}
	refreshToken, err := s.IssueRefreshToken(ctx, int(user.ID))
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, rawToken string) (*domain.AuthResponse, error) {
	tokenHash := hashToken(rawToken)

	rt, err := s.refreshTokenRepo.GetRefreshTokenByHash(ctx, tokenHash)
	if err != nil {
		return nil, domain.ErrInvalidRefreshToken
	}

	if rt.RevokedAt != nil || time.Now().After(rt.ExpiresAt) {
		return nil, domain.ErrInvalidRefreshToken
	}

	user, err := s.userRepo.GetByID(ctx, int64(rt.UserID))
	if err != nil {
		return nil, domain.ErrInvalidRefreshToken
	}

	if err := s.refreshTokenRepo.RevokeRefreshToken(ctx, rt.ID); err != nil {
		return nil, errors.New("error revoking old token")
	}

	accessToken, err := s.GenerateAccessToken(ctx, user)
	if err != nil {
		return nil, errors.New("error in generating token!")
	}

	newRefreshToken, err := s.IssueRefreshToken(ctx, int(user.ID))
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token:        accessToken,
		RefreshToken: newRefreshToken,
		User:         *user,
	}, nil
}
