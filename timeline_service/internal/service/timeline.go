package service

import (
	"github.com/tiltedEnmu/puregrade_timeline/internal/repository"
)

type TimelineService struct {
	repos repository.Repository
}

func (s *TimelineService) GetLatestById(userId, postId string, maxLen int64) ([]string, error) {
	index, err := s.repos.GetIndexByPostId(userId, postId, maxLen)
	if err != nil {
		return nil, err
	}
	return s.repos.GetRange(userId, 0, index)
}

func NewTimelineService(repos repository.Repository) *TimelineService {
	return &TimelineService{repos: repos}
}

func (s *TimelineService) GetRange(userId string, start, end int64) ([]string, error) {
	return s.repos.GetRange(userId, start, end)
}

func (s *TimelineService) Push(userId string, postIds ...string) (int64, error) {
	count, err := s.repos.Push(userId, postIds...)
	if err != nil {
		return 0, err
	}
	_, err = s.repos.Trim(userId, 3000)
	return count, err
}
