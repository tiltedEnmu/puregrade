package service

import (
	"github.com/tiltedEnmu/puregrade_user/internal/entities"
	"github.com/tiltedEnmu/puregrade_user/internal/repository"
)

type UserService struct {
	repos repository.Repository
}

func NewUserService(repos repository.Repository) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) CreateUser(user entities.User) (string, error) {
	if len(user.Roles) == 0 {
		user.Roles = append(user.Roles, 0)
	}

	id := "asd"

	err := s.repos.Create(user)

	return id, err
}

func (s *UserService) GetUser(id int64) (entities.User, error) {
	return s.repos.GetById(id)
}

func (s *UserService) FollowUser(id int64, publisherId int64) error {
	return s.repos.AddFollower(id, publisherId)
}

func (s *UserService) UnfollowUser(id int64, publisherId int64) error {
	return s.repos.DeleteFollower(id, publisherId)
}

func (s *UserService) Delete(id int64, password string) error {
	return s.repos.Delete(id, password)
}
