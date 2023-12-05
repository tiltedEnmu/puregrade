package repository

import (
	"github.com/tiltedEnmu/puregrade_user/internal/entities"
)

type Repository interface {
	// Receives user & save it in db
	Create(user entities.User) error
	Get(username string) (entities.User, error)
	GetById(id int64) (entities.User, error)
	AddFollower(id, publisherId int64) error
	DeleteFollower(id, publisherId int64) error
	Delete(id int64, password string) error
}
