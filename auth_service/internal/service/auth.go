package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/puregrade/puregrade-auth/internal/repository"
)

type AuthService struct {
	repos        repository.Repository
	jwtSecretKey string
}

func NewAuthService(repos repository.Repository, jwtSecretKey string) *AuthService {
	return &AuthService{
		repos:        repos,
		jwtSecretKey: jwtSecretKey,
	}
}

type jwtClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func (s *AuthService) ParseAccessToken(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
			}
			return []byte(s.jwtSecretKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(*jwtClaims); ok && parsedToken.Valid {
		return claims.UserId, nil
	}

	return "", errors.New("invalid access token")
}

func (s *AuthService) GenerateTokens(userId string) (string, string, error) {
	accessToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwtClaims{
			userId,
			jwt.StandardClaims{
				ExpiresAt: int64(30 * time.Minute),
			},
		},
	).SignedString([]byte(s.jwtSecretKey))

	if err != nil {
		return "", "", err
	}

	refreshToken := uuid.New().String()

	if err = s.repos.UpsertRefreshToken(userId, refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) GetUserId(refreshToken string) (string, error) {
	return s.repos.GetUserIdWithDelete(refreshToken)
}
