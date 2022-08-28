package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"flag"

	"github.com/skoflok/bandamper/config"
	storage "github.com/skoflok/bandamper/storage"
	"github.com/skoflok/bandamper/telegram"
	"github.com/skoflok/bandcamp_api_parser/api"
)

func main() {

	flag.Parse()
	command := flag.Arg(0)

	switch command {
	case "test-db":
		testDb()
	case "serve":
	case "releases":
		releases(flag.Args()[1:])
	case "fetch-first":
		fetchFirstRelease()
	case "fetch-page":
		fetchPage(flag.Args()[1:])
	case "migrate":
		migrate(checkSubcommand("Migration subcommand is not specified!"))
	case "run":
		run(checkSubcommand("Subcommand is not specified!"))
	case "tg":
		telegramCmd(flag.Args()[1:])
	case "help":
		help()
	case "ticker":
		ticker()
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

func fetchFirstRelease() {
	q := api.NewQueryArgs(0)
	r, err := api.FetchReleasesFromHome(q)
	if err != nil {
		log.Fatalf("Error from bandcamp api: %v", err)
	}

	if len(r.Items) > 0 {
		id, err := storage.StoreRelease(r.Items[0])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(id)
	}
}

func fetchPage(args []string) {
	page, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Page not defined %v", err)
	}
	q := api.NewQueryArgs(page)
	releases, err := api.FetchReleasesFromHome(q)
	if err != nil {
		log.Fatalf("Error from bandcamp api: %v", err)
	}

	count, err := storage.BulkStoreReleases(releases)
	if err != nil {
		log.Fatalf("Bulk store error: %v", err)
	}

	fmt.Printf("Store %d releases\n", count)
}

func releases(args []string) {
	input := "2006-01-02"
	sub := args[0]
	switch sub {
	case "get-curent-date":
		year, month, day := time.Now().Date()
		start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
		releases, err := storage.GetNotSentReleasesByDate(start, end)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(releases)
	case "get-by-date":
		date, err := time.Parse(input, args[1])
		if err != nil {
			log.Fatalln(err)
		}
		year, month, day := date.Date()
		start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
		releases, err := storage.GetNotSentReleasesByDate(start, end)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(releases)
	case "get-not-sent":
		year, month, day := time.Now().Date()
		start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
		releases, err := storage.GetNotSentReleasesByDate(start, end)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(releases)
		fmt.Println(releases[0].GetAlbumUrl())

	default:
		log.Fatalf("Comand %s no defined", sub)
	}
}

func telegramCmd(args []string) {
	sub := args[0]
	switch sub {
	case "publishing":

		year, month, day := time.Now().Date()
		end := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		start := end.Add(-1 * time.Hour * 24)

		releases, err := storage.GetNotSentReleasesByDate(start, end)
		if err != nil {
			log.Fatal(err)
		}

		for _, release := range releases {
			telegram.SendRelease(&release)
		}
	default:
		log.Fatalf("Comand %s no defined", sub)
	}
}

func ticker() {
	i := 0
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			i += 1
			fmt.Printf("Current time%d: %v\n", i, t)
		}
	}
}
