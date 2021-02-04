package model

import (
	"github.com/google/uuid"
	"time"
)

//Tag イベントに付随しているイベントの属性を説明するためのタグです
type Tag struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewTag 新しいTagを生成してポインタを返します
func NewTag(name string) *Tag {
	return &Tag{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
