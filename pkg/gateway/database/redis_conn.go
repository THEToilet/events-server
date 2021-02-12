package database

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewRedis() redis.Conn {
	//ip := "127.0.0.1"
	ip := "localhost"
	port := "6379"

	// コネクションプールの作成
	pool := newPool(ip + ":" + port)

	// コネクションの取得
	conn := pool.Get()

	return conn
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
