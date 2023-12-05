package service

import (
	_ "github.com/tiltedEnmu/puregrade_timeline/internal/entities"
	"github.com/tiltedEnmu/puregrade_timeline/internal/repository"
)

type Service interface {
	GetRange(userId string, start, end int64) ([]string, error)
	GetLatestById(userId, postId string, maxLen int64) ([]string, error)
	Push(userId string, postIds ...string) (int64, error)
}

func NewService(repos repository.Repository) Service {
	return NewTimelineService(repos)
}
