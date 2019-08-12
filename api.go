package goicqbot

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	client  *http.Client
	token   string
	baseURL string
	logger  *logrus.Logger
}

type Response struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
}

func (c *Client) Do(path string, params url.Values) ([]byte, error) {
	apiUrl, err := url.Parse(c.baseURL + path)
	params.Set("token", c.token)

	if err != nil {
		return []byte{}, fmt.Errorf("cannot parse url: %s", err)
	}

	apiUrl.RawQuery = params.Encode()
	req := &http.Request{
		URL: apiUrl,
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot make request to bot api: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("cannot close body: %s", err)
		}
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read body: %s", err)
	}

	c.logger.Debug(string(bytes))

	response := &Response{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return []byte{}, fmt.Errorf("cannot decode json: %s", err)
	}

	if !response.OK {
		return bytes, fmt.Errorf("error status from API: %s", response.Description)
	}

	return bytes, nil
}

func (c *Client) SendMessage(message Message) error {
	params := url.Values{
		"chatId":     []string{message.ChatID},
		"text":       []string{message.Text},
		"replyMsgId": []string{message.ReplyMsgID},
	}
	_, err := c.Do("/messages/sendText", params)

	return err
}

func (c *Client) GetEvents(lastEventID int, pollTime int) ([]*Event, error) {
	params := url.Values{
		"lastEventId": []string{strconv.Itoa(lastEventID)},
		"pollTime":    []string{strconv.Itoa(pollTime)},
	}
	events := EventsResponse{}

	body, err := c.Do("/events/get", params)
	if err != nil {
		return events.Events, fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(body, &events); err != nil {
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
