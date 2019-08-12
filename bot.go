package goicqbot

import (
	"context"
	"github.com/DmitryDorofeev/goicqbot/api"
)

const (
	NEW_MESSAGE_EVENT = "newMessage"
)

type Bot struct {
	ctx     context.Context
	client  *api.Client
	updater *api.Updater
}

func (b *Bot) SendMessage(chatID string, text string) error {
	return b.client.SendMessage(chatID, text)
}

func (b *Bot) GetUpdatesChannel() <-chan api.Event {
	updates := make(chan api.Event, 0)

	go b.updater.RunUpdatesCheck(updates)

	return updates
}

func NewBot(token string, opts ...BotOption) *Bot {
	apiUrl := "https://api.icq.net/bot/v1"

	for _, option := range opts {
		switch option.Type() {
		case "api_url":
			apiUrl = option.String()
		}
	}
	client := api.NewClient(apiUrl, token)
	updater := api.NewUpdater(client, 0)

	return &Bot{
		client:  client,
		updater: updater,
	}
}
