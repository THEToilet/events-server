package model

import (
	"github.com/google/uuid"
	"time"
)

//UserTag
type UserTag struct {
	UserTagID string
	EventID   string
	TagID     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewUserTag
func NewUserTag(eventID string, tagID string) *UserTag {
	return &UserTag{
		UserTagID: uuid.New().String(),
		EventID:   eventID,
		TagID:     tagID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
