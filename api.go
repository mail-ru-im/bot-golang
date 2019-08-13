package goicqbot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Client struct {
	client  *http.Client
	token   string
	baseURL string
	logger  *logrus.Logger
}

func (c *Client) Do(path string, params url.Values) ([]byte, error) {
	apiURL, err := url.Parse(c.baseURL + path)
	params.Set("token", c.token)

	if err != nil {
		return nil, fmt.Errorf("cannot parse url: %s", err)
	}

	apiURL.RawQuery = params.Encode()
	req := &http.Request{
		URL: apiURL,
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot make request to bot api: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("cannot close body: %s", err)
		}
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %s", err)
	}

	c.logger.Debug(string(bytes))

	response := &Response{}

	if err := response.UnmarshalJSON(bytes); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %s", err)
	}

	if !response.OK {
		return bytes, fmt.Errorf("error status from API: %s", response.Description)
	}

	return bytes, nil
}

func (c *Client) SendMessage(message Message) error {
	params := url.Values{
		"chatId":     {message.ChatID},
		"text":       {message.Text},
		"replyMsgId": {message.ReplyMsgID},
	}
	_, err := c.Do("/messages/sendText", params)

	return err
}

func (c *Client) GetEvents(lastEventID int, pollTime int) ([]*Event, error) {
	params := url.Values{
		"lastEventId": {strconv.Itoa(lastEventID)},
		"pollTime":    {strconv.Itoa(pollTime)},
	}
	events := EventsResponse{}

	body, err := c.Do("/events/get", params)
	if err != nil {
		return events.Events, fmt.Errorf("error while making request: %s", err)
	}

	if err := events.UnmarshalJSON(body); err != nil {
		return events.Events, fmt.Errorf("cannot parse events: %s", err)
	}

	return events.Events, nil
}

func NewClient(baseURL string, token string, logger *logrus.Logger) *Client {
	return &Client{
		token:   token,
		baseURL: baseURL,
		client:  http.DefaultClient,
		logger:  logger,
	}
}
