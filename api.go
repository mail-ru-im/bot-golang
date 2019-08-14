package goicqbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	c.logger.WithFields(logrus.Fields{
		"api_url": apiURL,
	}).Debug("requesting api")

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("request error")
		return []byte{}, fmt.Errorf("cannot make request to bot api: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("cannot close body")
		}
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("cannot read body")
		return []byte{}, fmt.Errorf("cannot read body: %s", err)
	}

	c.logger.WithFields(logrus.Fields{
		"response": string(bytes),
	}).Debug("got response from API")

	response := &Response{}

	if err := response.UnmarshalJSON(bytes); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %s", err)
	}

	if !response.OK {
		return bytes, fmt.Errorf("error status from API: %s", response.Description)
	}

	return bytes, nil
}

func (c *Client) GetInfo() (*BotInfo, error) {
	bytes, err := c.Do("/self/get", url.Values{})
	if err != nil {
		return nil, fmt.Errorf("error while receiving information: %s", err)
	}

	info := &BotInfo{}
	if err := info.UnmarshalJSON(bytes); err != nil {
		return nil, fmt.Errorf("error while unmarshalling information: %s", err)
	}

	return info, nil
}

func (c *Client) SendMessage(message *Message) error {
	params := url.Values{
		"chatId": []string{message.Chat.ID},
		"text":   []string{message.Text},
	}

	if message.ReplyMsgID != "" {
		params.Set("replyMsgId", message.ReplyMsgID)
	}

	if message.ForwardMsgID != "" {
		params.Set("forwardMsgId", message.ForwardMsgID)
		params.Set("forwardChatId", message.ForwardChatID)
	}

	bytes, err := c.Do("/messages/sendText", params)
	if err != nil {
		return fmt.Errorf("error while sending text: %s", err)
	}

	if err := json.Unmarshal(bytes, message); err != nil {
		return fmt.Errorf("cannot unmarshal response from API: %s", err)
	}

	return nil
}

func (c *Client) EditMessage(message *Message) error {
	params := url.Values{
		"msgId":  []string{message.ID},
		"chatId": []string{message.Chat.ID},
		"text":   []string{message.Text},
	}
	bytes, err := c.Do("/messages/editText", params)
	if err != nil {
		return fmt.Errorf("error while editing text: %s", err)
	}

	if err := json.Unmarshal(bytes, message); err != nil {
		return fmt.Errorf("cannot unmarshal response from API: %s", err)
	}

	return nil
}

func (c *Client) DeleteMessage(message *Message) error {
	params := url.Values{
		"msgId":  []string{message.ID},
		"chatId": []string{message.Chat.ID},
	}
	_, err := c.Do("/messages/deleteMessages", params)
	if err != nil {
		return fmt.Errorf("error while deleting message: %s", err)
	}

	return nil
}

func (c *Client) GetEvents(lastEventID int, pollTime int) ([]*Event, error) {
	params := url.Values{
		"lastEventId": {strconv.Itoa(lastEventID)},
		"pollTime":    {strconv.Itoa(pollTime)},
	}
	events := eventsResponse{}

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
