package database

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func Connect() redis.Conn {
	ip := "127.0.0.1"
	port := "6379"

	// redis-serverに接続する
	conn, err := redis.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Fatal("error")
	}
	defer conn.Close()
	return conn
}
