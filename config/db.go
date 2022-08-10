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
}

func NewDB() *db {
	db := &db{
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DBNAME"),
		os.Getenv("DATABASE_PROTOCOL"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
	}
	return db
}

func (db *db) String() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s", db.user, db.password, db.protocol, db.host, db.port, db.database)
}
