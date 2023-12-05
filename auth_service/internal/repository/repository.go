package repository

import (
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	// UpsertRefreshToken is
	UpsertRefreshToken(userId, token string) error
	GetUserIdWithDelete(token string) (string, error)
}

func NewRepository(db *redis.Client) Repository {
	return NewAuthRedis(db)
}
