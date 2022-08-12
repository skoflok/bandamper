package storage

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/skoflok/bandamper/config"
)

func Up() {
	db := Database(fmt.Sprint(config.NewDB()))
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql", driver)
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
