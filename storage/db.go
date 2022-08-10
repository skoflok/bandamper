package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Database(dsn string) *sql.DB {
	var db *sql.DB
	var err error
	if db, err = sql.Open("mysql", dsn); err != nil {
		log.Fatalf("Database connect error: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Database connect error: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	return db
}
