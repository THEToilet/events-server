package model

import "errors"

var (
	// ErrUserNotFound はユーザが存在しないエラーを表します。
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExisted はユーザが既に存在しているエラーを表します。
	ErrUserAlreadyExisted = errors.New("user has already existed")

	ErrSessionNotFound = errors.New("session not found")
	ErrSessionAlreadyExisted = errors.New("session has already existed")

	ErrTagNotFound = errors.New("tag not found")
	ErrTagAlreadyExisted = errors.New("tag has already existed")
)
