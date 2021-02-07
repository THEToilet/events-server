package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
)

type EventRepository interface {
	Find(id string) (*model.Event, error)
	FindAll() (*[]*model.Event, error)
	Save(id string) error
	Delete(id string) error
	Update(id string) error
}
