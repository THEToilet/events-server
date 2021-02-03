package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Driver名
const driverName = "mysql"

var Conn *gorm.DB

func Connect() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	// user:password@tcp(host:port)/database
	var err error
	Conn, err = gorm.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	return Conn
}
