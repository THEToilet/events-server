package model

import "github.com/google/uuid"

//User ログイン済みのユーザを表します
type User struct {
	ID   string
	Mail string
}

//NewUser 新しいUserを生成してポインタを返します
func NewUser(userMail string) *User {
	return &User{
		ID:   uuid.New().String(),
		Mail: userMail,
	}
}
