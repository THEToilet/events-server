package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Find(id uuid.UUID) (*model.User, error)
	FindAll() (*[]model.User, error)
	Save(id uuid.UUID, mail string) error
	Delete(id uuid.UUID) error
}
