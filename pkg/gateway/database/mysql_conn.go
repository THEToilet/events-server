package database

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/config"
	"log"
)

func NewMySqlDB() (*sql.DB, error) {
	conn, err := sql.Open("mysql", config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
