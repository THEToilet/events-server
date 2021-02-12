package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
)

type TagRepository interface {
	Save(tag model.Tag) error
	Find(id string) (*model.Tag, error)
	FindAll() ([]*model.Tag, error)
	Delete(id string) error
}
