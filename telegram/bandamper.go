package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendReleases() {
	_, _ = tgbotapi.NewBotAPI("MyAwesomeBotToken")

}
