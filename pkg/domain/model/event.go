package model

import (
	"github.com/google/uuid"
	"time"
)

//Event 通知をするイベント情報を表します。
type Event struct {
	ID          string
	PostedUser  string
	EventURL    string
	DeadLine    time.Time
	Description string
	Tag         []Tag
	CreatedAt   time.Time
	UpdateAt    time.Time
}

//NewEvent 新しいイベントを生成してポインタを返します
func NewEvent(postedUser string, eventURL string, deadLine time.Time, description string, tag []Tag) *Event {
	return &Event{
		ID: uuid.New().String(),
		PostedUser: postedUser,
		EventURL: eventURL,
		DeadLine: deadLine,
		Description: description,
		Tag: tag,
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}
