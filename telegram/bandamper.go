package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/skoflok/bandamper/config"
	"github.com/skoflok/bandamper/storage"
)

func SendReleases() {
	_, _ = tgbotapi.NewBotAPI("MyAwesomeBotToken")

}

func SendRelease(r *storage.Release) (err error) {
	tgConf := config.NewTg()
	bot, err := tgbotapi.NewBotAPI(tgConf.Token())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(tgConf.Channel(), r.ToTgMessage())

	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
	return
}
