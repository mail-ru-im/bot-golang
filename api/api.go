package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	client  http.Client
	token   string
	baseURL string
	debug   bool
}

type Response struct {
	OK bool `json:"ok"`
	Description string `json:"description"`
}

type Event struct {
	EventID int             `json:"eventId"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
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

	response := &Response{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return []byte{}, fmt.Errorf("cannot decode json: %s", err)
	}
	
	if !response.OK {
		return []byte{}, fmt.Errorf("error status from API: %s", response.Description)
	}

	return bytes, nil
}

func (c *Client) SendMessage(chatID string, text string) error {
	params := url.Values{
		"chatId": []string{chatID},
		"text":   []string{text},
	}
	_, err := c.Do("/messages/sendText", params)

	return err
}

func (c *Client) GetEvents(lastEventID int, pollTime int) ([]*Event, error) {
	params := url.Values{
		"lastEventId": []string{strconv.Itoa(lastEventID)},
		"pollTime":    []string{strconv.Itoa(pollTime)},
	}
	events := []*Event{}

	body, err := c.Do("/events/get", params)
	if err != nil {
		return events, fmt.Errorf("")
	}

	if err := json.Unmarshal(body, &events); err != nil {
		return events, fmt.Errorf("cannot parse events")
	}

	return events, nil
}

func NewClient(baseURL string, token string) *Client {
	return &Client{
		token:   token,
		baseURL: baseURL,
	}
}
