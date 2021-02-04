package repository

import (
	"github.com/google/uuid"
	"../model"
)

type UserRepository interface {
	find(id uuid.UUID) (*model.User, error)
	findAll() (*[]model.User, error)
	Save(id uuid.UUID, mail string) error
	delete(id uuid.UUID) error
}
