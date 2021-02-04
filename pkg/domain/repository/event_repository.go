package repository

import (
	"github.com/google/uuid"
	"../model"
)

type EventRepository interface {
	Find(id uuid.UUID) (*model.Tag, error)
	FindAll() (*[]model.Event, error)
	Save(id uuid.UUID, mail string) error
	Delete(id uuid.UUID) error
}
