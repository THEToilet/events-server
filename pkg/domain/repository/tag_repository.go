package repository

import (
	"github.com/google/uuid"
	"../model"
)

type TagRepository interface {
	Save(name string) (*model.Tag, error)
	Find(id uuid.UUID) (*model.Tag, error)
	FindAll() (*[]model.Tag, error)
	Delete(id uuid.UUID) error
}

