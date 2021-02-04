package gateway

import (
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

func Set(userId string) (string, error) {
	session, err := uuid.NewUUID()
	s := session.String()
	redis.Connect().Do("SET", userId, s)
	return s, err
}

func Get(userId string) (string, error) {
	s, err := redis.String(conn.Do("GET", "temperature"))
	if err != nil {
		panic(err)
	}

	return s, err
}
