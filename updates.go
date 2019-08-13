package goicqbot

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type Updater struct {
	logger      *logrus.Logger
	client      *Client
	lastEventID int
	PollTime    int
}

func (u *Updater) RunUpdatesCheck(ctx context.Context, ch chan<- Event) {
	_, err := u.GetLastEvents(0)
	if err != nil {
		log.Printf("cannot make initial request to events: %s", err)
	}

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			events, err := u.GetLastEvents(u.PollTime)
			if err != nil {
				u.logger.Info(err)
				u.logger.Info("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, event := range events {
				ch <- *event
			}
		}
	}
}

func (u *Updater) GetLastEvents(pollTime int) ([]*Event, error) {
	events, err := u.client.GetEvents(u.lastEventID, pollTime)
	if err != nil {
		u.logger.Debug(events)
		return events, fmt.Errorf("cannot get events: %s", err)
	}

	count := len(events)
	if count > 0 {
		u.lastEventID = events[count-1].EventID
	}

	return events, nil
}

func NewUpdater(client *Client, pollTime int, logger *logrus.Logger) *Updater {
	if pollTime == 0 {
		pollTime = 60
	}

	return &Updater{
		client:      client,
		lastEventID: 0,
		PollTime:    pollTime,
		logger:      logger,
	}
}
