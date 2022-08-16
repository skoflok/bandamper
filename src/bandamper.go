package main

import (
	"fmt"
	"log"
	"os"

	"flag"

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
	os.Setenv("DATABASE_DRIVER", "mysql")
	// fmt.Println(parser.FetchReleasesFromHome(parser.NewQueryArgs(0)))

	flag.Parse()
	command := flag.Arg(0)
	fmt.Println(command)

	switch command {
	case "test-db":
		testDb()
	case "serve":
	case "migrate":
		migrate(checkSubcommand("Migration subcommand is not specified!"))
	case "run":
		run(checkSubcommand("Subcommand is not specified!"))
	case "help":
		help()
	default:
		help()
	}
}

func help() {
	fmt.Println("Some help string")
}

func testDb() {
	cfg := config.NewDB()
	db := storage.Open(cfg.Driver(), fmt.Sprint(cfg))
	defer db.Close()

	fmt.Println(cfg)
	fmt.Println(db)
}

func run(cmd string) (err error) {
	return
}

func migrate(cmd string) (err error) {
	switch cmd {
	case "up":
		storage.Up()
	case "down":
		storage.Down()
	case "drop":
		storage.Drop()
	default:
		log.Fatalf("Unexpected migrate command %s", cmd)
	}
	return
}

func checkSubcommand(callbackError string) (command string) {
	command = flag.Arg(1)
	if command == "" {
		log.Fatalln(callbackError)
	}
	return command
}
