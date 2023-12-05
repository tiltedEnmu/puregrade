package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisTimeline struct {
	client *redis.Client
}

func NewRedisTimeline(client *redis.Client) *RedisTimeline {
	return &RedisTimeline{client: client}
}

func (r *RedisTimeline) Push(userId string, postIds ...string) (int64, error) {
	// []string to []interface
	s := make([]interface{}, len(postIds))
	for i, v := range postIds {
		s[i] = v
	}

	return r.client.LPush(context.Background(), "timeline:"+userId, s...).Result()
}

func (r *RedisTimeline) GetRange(userId string, start, end int64) ([]string, error) {
	return r.client.LRange(context.Background(), "timeline:"+userId, start, end).Result()
}

func (r *RedisTimeline) GetIndexByPostId(userId, postId string, maxLen int64) (int64, error) {
	return r.client.LPos(
		context.Background(),
		"timeline:"+userId,
		postId,
		redis.LPosArgs{
			Rank:   1,
			MaxLen: maxLen,
		},
	).Result()
}

func (r *RedisTimeline) Trim(userId string, cap int32) (string, error) {
	return r.client.LTrim(context.Background(), "timeline:"+userId, 0, int64(cap-1)).Result()
}
