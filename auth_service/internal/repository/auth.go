package repository

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type AuthRedis struct {
	db *redis.Client
}

func NewAuthRedis(db *redis.Client) *AuthRedis {
	return &AuthRedis{db: db}
}

func (r *AuthRedis) UpsertRefreshToken(userId, token string) error {
	res := r.db.HSet(context.Background(), "refresh_tokens", token, userId)
	log.Print(res)
	return res.Err()
}

func (r *AuthRedis) GetUserIdWithDelete(token string) (string, error) {
	id, err := r.db.HGet(context.Background(), "refresh_tokens", token).Result()
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	r.db.HDel(context.Background(), "refresh_tokens", token)
	return id, err
}
