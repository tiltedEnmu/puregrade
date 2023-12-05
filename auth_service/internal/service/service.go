package service

import (
	"github.com/puregrade/puregrade-auth/internal/repository"
)

type Service interface {
	GenerateTokens(userId string) (string, string, error) // access, refresh, err
	ParseAccessToken(token string) (string, error)
	GetUserId(token string) (string, error)
}

func NewService(repos repository.Repository, jwtSecretKey string) Service {
	return NewAuthService(repos, jwtSecretKey)
}
