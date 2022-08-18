package config

import (
	"log"
	"os"
	"strconv"
)

type Tg struct {
	token   string
	channel string
}

func NewTg() *Tg {
	return &Tg{
		os.Getenv("BOT_TOKEN"),
		os.Getenv("CHANNEL_ID"),
	}
}

func (t *Tg) Token() string {
	return t.token
}

func (t *Tg) Channel() int64 {
	id, err := strconv.ParseInt(t.channel, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
