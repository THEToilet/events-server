package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/google/uuid"
)

type TagRepository interface {
	Save(name string) (*model.Tag, error)
	Find(id uuid.UUID) (*model.Tag, error)
	FindAll() (*[]model.Tag, error)
	Delete(id uuid.UUID) error
}
