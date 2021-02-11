package model

import "github.com/google/uuid"

//User ログイン済みのユーザを表します
type User struct {
	UserID       string
	UserMail     string
	UserPassword string
}

//NewUser 新しいUserを生成してポインタを返します
func NewUser(userMail string, userPassword string) *User {
	return &User{
		UserID:   uuid.New().String(),
		UserMail: userMail,
		UserPassword: userPassword,
	}
}
