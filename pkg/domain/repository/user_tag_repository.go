package repository

import (
"github.com/THEToilet/events-server/pkg/domain/model"
)

type UserTagRepository interface {
	Save(name string) (*model.UserTag, error)
	Find(id string) (*model.UserTag, error)
	FindAll() ([]*model.UserTag, error)
	Delete(id string) error
}