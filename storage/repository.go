package storage

import (
	"fmt"
	"time"

	"github.com/skoflok/bandamper/config"
	"github.com/skoflok/bandcamp_api_parser/api"
)

func StoreRelease(r api.Release) (rowId int64, err error) {
	dbConf := config.NewDB()

	db := Open(dbConf.Driver(), dbConf.String())
	defer db.Close()

	stmt, err := db.Prepare(getInsertReleaseQuery())

	if err != nil {
		return 0, fmt.Errorf("Statement prepare error: %v", err)
	}

	defer stmt.Close()

	layout := "_2 Jan 2006 03:04:05 GMT"

	publishDate, err := time.Parse(layout, r.PublishDate)

	if err != nil {
		return 0, fmt.Errorf("Publish date parse error: %v", err)
	}

	result, err := stmt.Exec(
		r.Type,
		r.Id,
		r.BandId,
		r.IsPreorder,
		publishDate.Format("2006-01-02 03:04:05"),
		r.Genre,
		r.Album,
		r.Artist,
		r.FeaturedTrack.File.Link,
		r.UrlHints.Subdomain,
		r.UrlHints.Slug,
	)
	if err != nil {
		return 0, fmt.Errorf("Exec query error: %v", err)
	}
	return result.LastInsertId()
}

func getReleaseByDate(start time.Time, end time.Time) (items []string, err error) {
	return
}
