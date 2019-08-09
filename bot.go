package goicqbot

import (
	"context"
	"github.com/DmitryDorofeev/goicqbot/api"
	"log"
	"time"
)

type Bot struct {
	ctx     context.Context
	client  *api.Client
	updater *api.Updater
}

type Update struct {
	Type    string
	Payload map[string]string
}

func (b *Bot) SendMessage(chatID string, text string) error {
	return b.client.SendMessage(chatID, text)
}

func (b *Bot) runUpdatesCheck(ch chan<- Update) {
	for {
		updates, err := b.updater.GetUpdates()
		if err != nil {
			log.Println(err)
			log.Println("Failed to get updates, retrying in 3 seconds...")
			time.Sleep(time.Second * 3)

			continue
		}

		for _, update := range updates {
			ch <- Update{Type: update.Type}
		}
	}
}

func (b *Bot) GetUpdatesChannel() <-chan Update {
	updates := make(chan Update, 0)

	go b.runUpdatesCheck(updates)

	return updates
}

func NewBot(token string) *Bot {
	client := api.NewClient("https://api.icq.net/bot/v1", token)
	updater := api.NewUpdater(client, 0)

	return &Bot{
		client:  client,
		updater: updater,
	}
}
