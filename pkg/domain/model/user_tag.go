package model

import (
	"time"
)

//UserTag
type UserTag struct {
	EventID   string
	TagID     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewUserTag
func NewUserTag(eventID string, tagID string) *UserTag {
	return &UserTag{
		EventID:   eventID,
		TagID:     tagID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
