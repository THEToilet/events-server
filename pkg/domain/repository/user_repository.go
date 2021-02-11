package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
)

type UserRepository interface {
	Find(id string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Save(id string, mail string) error
	Delete(id string) error
}
