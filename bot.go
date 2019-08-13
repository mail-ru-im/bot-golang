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

// SendMessage sends a message, passed as an argument
// This method fills the argument with ID of sent message and returns an error if any
func (b *Bot) SendMessage(message *Message) error {
	return b.client.SendMessage(message)
}

// EditMessage edit a message passed as an argument
func (b *Bot) EditMessage(message *Message) error {
	return b.client.EditMessage(message)
}

// GetUpdatesChannel returns a channel, which will be filled with events
// You can pass cancellable context there and stop receiving events
// The channel will be closed after context cancellation
func (b *Bot) GetUpdatesChannel(ctx context.Context) <-chan Event {
	updates := make(chan Event, 0)

	go b.updater.RunUpdatesCheck(ctx, updates)

	return updates
}

// NewBot returns new bot object
// All communications with ICQ bot API must go through Bot struct
func NewBot(token string, opts ...BotOption) *Bot {
	debug := false
	apiUrl := "https://api.icq.net/bot/v1"
	logger := logrus.New()

	for _, option := range opts {
		switch option.Type() {
		case "api_url":
			apiUrl = option.String()
		case "debug":
			debug = option.Bool()
		}
	}

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	client := NewClient(apiUrl, token, logger)
	updater := NewUpdater(client, 0, logger)

	return &Bot{
		client:  client,
		updater: updater,
		logger:  logger,
	}
}
