package database

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewMySqlDB() (*sql.DB, error) {
	conn, err := sql.Open("mysql", config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
