package service

import (
	"cryptoMonitor/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	API *tgbotapi.BotAPI
}

var telegramBot *TelegramBot

func InitTelegramBot() (err error) {
	bot, err := tgbotapi.NewBotAPI(config.Get().TelegramBot.APIToken)
	if err != nil {
		return
	}
	telegramBot = &TelegramBot{
		API: bot,
	}

	return
}

func GetTelegramBot() *TelegramBot {
	return telegramBot
}

func (t *TelegramBot) SendMessage(message string) (err error) {
	msg := tgbotapi.NewMessage(config.Get().TelegramBot.ChatId, message)
	_, err = t.API.Send(msg)
	if err != nil {
		return
	}

	return
}
