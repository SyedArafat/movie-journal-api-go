package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

// DB holds the database connection pool
type DB struct {
	SQL *gorm.DB
}

var dbConn = &DB{}

var maxOpenDbConn, _ = strconv.Atoi(os.Getenv("MAX_OPEN_DB_CONNECTIONS"))
var maxIdleDbConn, _ = strconv.Atoi(os.Getenv("MAX_IDLE_DB_CONNECTIONS"))
var dbLifeTimeInMinute, _ = strconv.Atoi(os.Getenv("MAX_DB_LIFETIME_IN_MINUTE"))
var maxDbLifeTime = time.Duration(dbLifeTimeInMinute) * time.Minute

func ConnectSQL() (*DB, error) {
	d, err := NewDatabase(dsn())
	if err != nil {
		//logger.Error.Println("DB connection failed", err)
		panic(err)
	}

	db, _ := d.DB()
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d
	err = testDB(dbConn.SQL)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testDB(gdb *gorm.DB) error {
	db, _ := gdb.DB()
	err := db.Ping()
	if err != nil {
		//logger.Error.Println("Database failed", err)
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
