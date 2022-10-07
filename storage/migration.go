package storage

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/skoflok/bandamper/config"
	"github.com/skoflok/bandamper/helpers"
)

func initMigrate() *migrate.Migrate {
	dbConf := config.NewDB()
	db := Open(dbConf.Driver(), dbConf.String())
	defer db.Close()

	instance, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Migration instance error: %v", err)
	}

	wd := helpers.Wd()

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/migrations/%s", wd, dbConf.Driver()),
		dbConf.Driver(),
		instance,
	)

	if err != nil {
		log.Fatalf("Migration db istance error: %v", err)
	}
	return m
}

func Up() {
	m := initMigrate()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration Up error: %v", err)
	}
}

func Down() {
	m := initMigrate()
	if err := m.Down(); err != nil {
		log.Fatalf("Migration Down error: %v", err)
	}
}

func Drop() {
	m := initMigrate()
	if err := m.Drop(); err != nil {
		log.Fatalf("Migration Drop error: %v", err)
	}
}
