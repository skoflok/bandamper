package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/skoflok/bandamper/config"
	"github.com/skoflok/bandamper/storage"
)

type botApi struct {
	Bot       *tgbotapi.BotAPI
	ChannelId int64
}

func newTgBot() botApi {
	tgConf := config.NewTg()
	newbot, err := tgbotapi.NewBotAPI(tgConf.Token())
	if err != nil {
		log.Fatalf("New bot API failed: %v", err)
	}
	botApi := botApi{
		Bot:       newbot,
		ChannelId: tgConf.Channel(),
	}
	return botApi
}

func SendReleases() {
	_, _ = tgbotapi.NewBotAPI("MyAwesomeBotToken")

}

func SendRelease(r *storage.Release) (err error) {
	botApi := newTgBot()

	msg := tgbotapi.NewMessage(botApi.ChannelId, r.ToTgMessage())
	msg.ParseMode = "Markdown"

	if _, err := botApi.Bot.Send(msg); err != nil {
		panic(err)
	} else {
		r.SetSendingStatus(true)
	}
	return
}
