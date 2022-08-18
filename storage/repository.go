package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/skoflok/bandamper/config"
	"github.com/skoflok/bandcamp_api_parser/api"
)

type Release struct {
	Id            int64
	Type          string
	ReleaseId     int
	BandId        int
	IsPreorder    bool
	PublishDate   time.Time
	Genre         string
	Album         string
	Artist        string
	FeaturedTrack string
	Subdomain     string
	Slug          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

const apiLayout = "_2 Jan 2006 15:04:05 GMT"
const dbLayout = "2006-01-02 15:04:05"

func BulkStoreReleases(releases api.Releases) (count int, err error) {
	if len(releases.Items) == 0 {
		return 0, fmt.Errorf("Empty Releases Items")
	}

	for _, r := range releases.Items {
		if _, err := StoreRelease(r); err != nil {
			return count, fmt.Errorf("Bulk store releases (release_id %d) error: %v", r.Id, err)
		} else {
			count += 1
		}
	}
	return
}

func StoreRelease(r api.Release) (rowId int64, err error) {

	exist, ok := GetReleaseByReleaseId(r.Id)
	if !ok {
		return insertRelease(r)
	} else {
		return updateRelease(r, exist.Id)
	}
}

func GetReleaseByDate(start time.Time, end time.Time) (items []string, err error) {
	return
}

func GetReleaseByReleaseId(id int) (r *Release, ok bool) {

	dbConf := config.NewDB()
	db := Open(dbConf.Driver(), dbConf.String())
	defer db.Close()

	query := getReleaseByIdQuery()

	row := db.QueryRow(query, id)
	r = &Release{}
	if err := row.Scan(
		&r.Id,
		&r.Type,
		&r.ReleaseId,
		&r.BandId,
		&r.IsPreorder,
		&r.PublishDate,
		&r.Genre,
		&r.Album,
		&r.Artist,
		&r.FeaturedTrack,
		&r.Subdomain,
		&r.Slug,
		&r.UpdatedAt,
		&r.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return r, false
		} else {
			log.Fatalf("Database select error : %v", err)
		}
	}
	return r, true
}

func (r *Release) GetAlbumUrl() (url string, ok bool) {
	return
}

func insertRelease(r api.Release) (rowId int64, err error) {

	dbConf := config.NewDB()
	db := Open(dbConf.Driver(), dbConf.String())
	defer db.Close()

	stmt, err := db.Prepare(getInsertReleaseQuery())

	if err != nil {
		return 0, fmt.Errorf("Insert statement prepare error: %v", err)
	}

	defer stmt.Close()

	publishDate, err := time.Parse(apiLayout, r.PublishDate)
	updatedAt := time.Now()
	createdAt := time.Now()

	if err != nil {
		return 0, fmt.Errorf("Publish date parse error: %v", err)
	}

	result, err := stmt.Exec(
		r.Type,
		r.Id,
		r.BandId,
		r.IsPreorder,
		publishDate.Format(dbLayout),
		r.Genre,
		r.Album,
		r.Artist,
		r.FeaturedTrack.File.Link,
		r.UrlHints.Subdomain,
		r.UrlHints.Slug,
		updatedAt.Format(dbLayout),
		createdAt.Format(dbLayout),
	)
	if err != nil {
		return 0, fmt.Errorf("Exec query error: %v", err)
	}
	return result.LastInsertId()
}

func updateRelease(r api.Release, id int64) (rowId int64, err error) {
	dbConf := config.NewDB()
	db := Open(dbConf.Driver(), dbConf.String())
	defer db.Close()

	stmt, err := db.Prepare(getUpdateReleaseByIdQuery())

	if err != nil {
		return 0, fmt.Errorf("Update statement prepare error: %v", err)
	}

	defer stmt.Close()

	publishDate, err := time.Parse(apiLayout, r.PublishDate)
	updatedAt := time.Now()

	result, err := stmt.Exec(
		r.Type,
		r.BandId,
		r.IsPreorder,
		publishDate.Format(dbLayout),
		r.Genre,
		r.Album,
		r.Artist,
		r.FeaturedTrack.File.Link,
		r.UrlHints.Subdomain,
		r.UrlHints.Slug,
		updatedAt.Format(dbLayout),
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("Exec update query error: %v", err)
	}
	_, err = result.LastInsertId()

	return id, err
}
