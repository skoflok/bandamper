package main

import (
	"fmt"
	"os"

	"github.com/skoflok/bandamper/config"
	storage "github.com/skoflok/bandamper/storage"
	_ "github.com/skoflok/bandcamp_api_parser/api"
)

func main() {
	os.Setenv("DATABASE_USER", "localwp")
	os.Setenv("DATABASE_PASSWORD", "localwp")
	os.Setenv("DATABASE_DBNAME", "bandcamp")
	os.Setenv("DATABASE_PROTOCOL", "tcp")
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "3306")
	// fmt.Println(parser.FetchReleasesFromHome(parser.NewQueryArgs(0)))
	fmt.Println(config.NewDB())
	db := storage.Open(fmt.Sprint(config.NewDB()))
	defer db.Close()
	fmt.Println(db)
}
