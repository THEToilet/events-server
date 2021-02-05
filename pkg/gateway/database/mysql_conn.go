package database

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/config"
	"log"
)

func NewMySqlDB() (*sql.DB, error) {
	/*
		user := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		database := os.Getenv("MYSQL_DATABASE")
	*/
	//fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	// user:password@tcp(host:port)/database
	var err error
	conn, err := sql.Open("mysql", config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
