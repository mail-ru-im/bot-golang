package api

type Updater struct {
	client *Client
	lastEventID int
	PollTime int
}

func (u *Updater) GetUpdates() ([]*Event, error) {
	_, err := u.client.GetEvents(u.lastEventID, u.PollTime)

	return []*Event{}, err
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

