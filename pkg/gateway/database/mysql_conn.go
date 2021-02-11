package database

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/config"
	"github.com/THEToilet/events-server/pkg/log"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDB() (*sql.DB, error) {
	logger := log.New()
	conn, err := sql.Open("mysql", config.DSN())
	if err != nil {
		logger.Error("mysql connection refused")
		return nil, err
	}
	return conn, err
}
