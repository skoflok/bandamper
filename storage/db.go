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
	return "INSERT INTO releases(type, release_id, band_id, is_preorder, publish_date, genre, album, artist, featured_track, subdomain, slug, updated_at, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
}

func getReleaseByIdQuery() string {
	return `SELECT
id,
type,
release_id,
band_id,
is_preorder,
publish_date,
genre,
album,
artist,
featured_track,
subdomain,
slug,
updated_at,
created_at,
is_sent
FROM releases where release_id = ?
;`
}

func getUpdateReleaseByIdQuery() string {
	return `UPDATE releases set 
	type = ?,
	band_id = ?,
	is_preorder = ?,
	publish_date = ? ,
	genre = ? ,
	album = ? ,
	artist = ? ,
	featured_track = ?,
	subdomain = ?,
	slug = ? ,
	updated_at = ?,
	is_sent = ?
	where id = ?`
}

func getNotSentReleasesByDateQuery() string {
	return `SELECT
	id,
	type,
	release_id,
	band_id,
	is_preorder,
	publish_date,
	genre,
	album,
	artist,
	featured_track,
	subdomain,
	slug,
	updated_at,
	created_at
	FROM releases 
	where is_sent = 0 and is_preorder = 0
	and publish_date BETWEEN ? AND ?
	;`
}

func getUpdateSendingStatusQuery() string {
	return `UPDATE releases set 
	is_sent = ?
	where id = ?`
}
