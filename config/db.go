package config

import (
	"fmt"
	"os"
)

type db struct {
	user     string
	password string
	database string
	protocol string
	host     string
	port     string
	driver   string
}

func NewDB() *db {
	db := &db{
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DBNAME"),
		os.Getenv("DATABASE_PROTOCOL"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_DRIVER"),
	}
	return db
}

func (db *db) String() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?multiStatements=true", db.user, db.password, db.protocol, db.host, db.port, db.database)
}

func NewDSN() string {
	return fmt.Sprint(NewDB())
}

func (db *db) Driver() string {
	return db.driver
}
