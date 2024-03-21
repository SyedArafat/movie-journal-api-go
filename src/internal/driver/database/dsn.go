package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Connection struct {
	Username string
	Password string
	Hostname string
	Dbname   string
	Port     string
}

func getConnectionString() Connection {
	return Connection{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Hostname: os.Getenv("DB_HOST"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_HOST_PORT"),
	}
}

func dsn() string {
	connection := getConnectionString()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection.Username, connection.Password, connection.Hostname, connection.Port, connection.Dbname)
}
