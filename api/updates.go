package api

import (
	"log"
	"time"
)

type Updater struct {
	client      *Client
	lastEventID int
	PollTime    int
}

func (u *Updater) RunUpdatesCheck(ch chan<- Event) {
	for {
		events, err := u.GetLastEvents()
		if err != nil {
			log.Println(err)
			log.Println("Failed to get updates, retrying in 3 seconds...")
			time.Sleep(time.Second * 3)

			continue
		}

		for _, event := range events {
			ch <- *event
		}
	}
}

func (u *Updater) GetLastEvents() ([]*Event, error) {
	events, err := u.client.GetEvents(u.lastEventID, u.PollTime)
	if err != nil {
		return events, nil
	}

	count := len(events)
	if count > 0 {
		u.lastEventID = events[count-1].EventID
	}

	return events, nil
}

func NewUpdater(client *Client, pollTime int) *Updater {
	if pollTime == 0 {
		pollTime = 60
	}

	return &Updater{
		client:      client,
		lastEventID: 0,
		PollTime:    pollTime,
	}
}
