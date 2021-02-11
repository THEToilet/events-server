package model

import (
	"github.com/google/uuid"
	"time"
)

//Event 通知をするイベント情報を表します。
type Event struct {
	EventID     string
	PostedUser  string
	EventURL    string
	DeadLine    time.Time
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//NewEvent 新しいイベントを生成してポインタを返します
func NewEvent(postedUser string, eventURL string, deadLine time.Time, description string) *Event {
	return &Event{
		EventID:     uuid.New().String(),
		PostedUser:  postedUser,
		EventURL:    eventURL,
		DeadLine:    deadLine,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
