package goicqbot

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	NEW_MESSAGE_EVENT = "newMessage"
)

type Bot struct {
	ctx     context.Context
	client  *Client
	updater *Updater
	logger  *logrus.Logger
}

func (b *Bot) SendMessage(message Message) error {
	return b.client.SendMessage(message)
}

func (b *Bot) GetUpdatesChannel() <-chan Event {
	updates := make(chan Event, 0)

	go b.updater.RunUpdatesCheck(updates)

	return updates
}

func NewBot(token string, opts ...BotOption) *Bot {
	debug := false
	apiURL := "https://api.icq.net/bot/v1"
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	for _, option := range opts {
		switch option.Type() {
		case "api_url":
			apiURL = option.String()
		case "debug":
			debug = option.Bool()
		}
	}

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	client := NewClient(apiURL, token, logger)
	updater := NewUpdater(client, 0, logger)

	return &Bot{
		client:  client,
		updater: updater,
		logger:  logger,
	}
}
