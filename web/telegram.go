package web

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

type TelegramClient struct {
	Client *tele.Bot
	ChatID tele.ChatID
}

func NewTelegramClient(token string, chatID string) *TelegramClient {
	pref := tele.Settings{
		Token: token,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(fmt.Errorf("could not instanciate Telegram client: %s", err))
	}

	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		log.Fatal(fmt.Errorf("could not convert chatID %s to int64: %s", chatID, err))
	}

	return &TelegramClient{
		Client: b,
		ChatID: tele.ChatID(i),
	}
}
