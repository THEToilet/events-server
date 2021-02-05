package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
)

type EventRepository interface {
	Find(id string) (*model.Tag, error)
	FindAll() (*[]model.Event, error)
	Save(id string, mail string) error
	Delete(id string) error
}
