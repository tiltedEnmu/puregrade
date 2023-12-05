package service

import (
	"github.com/tiltedEnmu/puregrade_user/internal/entities"
	"github.com/tiltedEnmu/puregrade_user/internal/repository"
)

type Service interface {
	CreateUser(user entities.User) (int64, error)
	GetUser(id int64) (entities.User, error)
	FollowUser(id int64, publisherId int64) error
	UnfollowUser(id int64, publisherId int64) error
	Delete(id int64, password string) error
}

func NewService(repos repository.Repository) Service {
	return NewUserService(repos)
}
