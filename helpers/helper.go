package helpers

import (
	"log"
	"os"
)

func Wd() (wd string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not retrieve current working directory: %v", err)
	}
	return wd
}
