package model

import (
	"github.com/google/uuid"
	"time"
)

//Tag イベントに付随しているイベントの属性を説明するためのタグです
type Tag struct {
	TagID     string
	TagName   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewTag 新しいTagを生成してポインタを返します
func NewTag(name string) *Tag {
	return &Tag{
		TagID:     uuid.New().String(),
		TagName:   name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
