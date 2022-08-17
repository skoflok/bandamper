package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Open(driver string, dsn string) *sql.DB {
	var db *sql.DB
	var err error
	if db, err = sql.Open(driver, dsn); err != nil {
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

func getInsertReleaseQuery() string {
	return "INSERT INTO releases(type, release_id, band_id, is_preorder, publish_date, genre, album, artist, featured_track, subdomain, slug) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}
