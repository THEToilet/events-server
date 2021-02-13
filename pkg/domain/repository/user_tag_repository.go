package repository

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
)

type UserTagRepository interface {
	Save(*model.UserTag) error
	FindAll(eventID string) ([]*model.UserTag, error)
	Delete(string, string) error
}
