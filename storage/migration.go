package storage

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up() {
	db := Database()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql", driver)
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
