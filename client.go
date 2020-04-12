package botgolang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Client struct {
	client  *http.Client
	token   string
	baseURL string
	logger  *logrus.Logger
}

func (c *Client) Do(path string, params url.Values, file *os.File) ([]byte, error) {
	apiURL, err := url.Parse(c.baseURL + path)
	params.Set("token", c.token)

	if err != nil {
		return nil, fmt.Errorf("cannot parse url: %s", err)
	}

	apiURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil || req == nil {
		return nil, fmt.Errorf("cannot init http request: %s", err)
	}

	if file != nil {
		buffer := &bytes.Buffer{}
		multipartWriter := multipart.NewWriter(buffer)

		fileWriter, err := multipartWriter.CreateFormFile("file", file.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot create multipart writer: %s", err)
		}

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, fmt.Errorf("cannot copy file into buffer: %s", err)
		}

		if err := multipartWriter.Close(); err != nil {
			return nil, fmt.Errorf("cannot close multipartWriter: %s", err)
		}

		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		req.Body = ioutil.NopCloser(buffer)
		req.Method = http.MethodPost
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

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("cannot read body")
		return []byte{}, fmt.Errorf("cannot read body: %s", err)
	}

	c.logger.WithFields(logrus.Fields{
		"response": string(responseBody),
	}).Debug("got response from API")

	response := &Response{}

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %s", err)
	}

	if !response.OK {
		return responseBody, fmt.Errorf("error status from API: %s", response.Description)
	}

	return responseBody, nil
}

func (c *Client) GetInfo() (*BotInfo, error) {
	response, err := c.Do("/self/get", url.Values{}, nil)
	if err != nil {
		return nil, fmt.Errorf("error while receiving information: %s", err)
	}

	info := &BotInfo{}
	if err := json.Unmarshal(response, info); err != nil {
		return nil, fmt.Errorf("error while unmarshalling information: %s", err)
	}

	return info, nil
}

func (c *Client) GetChatInfo(chatID string) (*Chat, error) {
	params := url.Values{
		"chatId": {chatID},
	}
	response, err := c.Do("/chats/getInfo", params, nil)
	if err != nil {
		return nil, fmt.Errorf("error while receiving information: %s", err)
	}

	chat := &Chat{}
	if err := json.Unmarshal(response, chat); err != nil {
		return nil, fmt.Errorf("error while unmarshalling information: %s", err)
	}

	if chat.Group != "group" {
		return chat, nil
	}

	response, err = c.Do("/chats/getAdmins", params, nil)
	if err != nil {
		return nil, fmt.Errorf("error while receiving admins: %s", err)
	}

	if err := json.Unmarshal(response, chat); err != nil {
		return nil, fmt.Errorf("error while unmarshalling admins: %s", err)
	}

	return chat, nil
}

func (c *Client) GetFileInfo(fileID string) (*File, error) {
	params := url.Values{
		"fileId": {fileID},
	}
	response, err := c.Do("/files/getInfo", params, nil)
	if err != nil {
		return nil, fmt.Errorf("error while receiving information: %s", err)
	}

	file := &File{}
	if err := json.Unmarshal(response, file); err != nil {
		return nil, fmt.Errorf("error while unmarshalling information: %s", err)
	}

	return file, nil
}

func (c *Client) GetVoiceInfo(fileID string) (*File, error) {
	return c.GetFileInfo(fileID)
}

func (c *Client) SendTextMessage(message *Message) error {
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

	if message.InlineKeyboard != nil {
		data, err := json.Marshal(message.InlineKeyboard)
		if err != nil {
			return fmt.Errorf("cannot marshal inline keyboard markup: %s", err)
		}

		params.Set("inlineKeyboardMarkup", string(data))
	}

	response, err := c.Do("/messages/sendText", params, nil)
	if err != nil {
		return fmt.Errorf("error while sending text: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
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
	response, err := c.Do("/messages/editText", params, nil)
	if err != nil {
		return fmt.Errorf("error while editing text: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
		return fmt.Errorf("cannot unmarshal response from API: %s", err)
	}

	return nil
}

func (c *Client) DeleteMessage(message *Message) error {
	params := url.Values{
		"msgId":  []string{message.ID},
		"chatId": []string{message.Chat.ID},
	}
	_, err := c.Do("/messages/deleteMessages", params, nil)
	if err != nil {
		return fmt.Errorf("error while deleting message: %s", err)
	}

	return nil
}

func (c *Client) SendFileMessage(message *Message) error {
	params := url.Values{
		"chatId":  {message.Chat.ID},
		"caption": {message.Text},
		"fileId":  {message.FileID},
	}

	if message.ReplyMsgID != "" {
		params.Set("replyMsgId", message.ReplyMsgID)
	}

	if message.ForwardMsgID != "" {
		params.Set("forwardMsgId", message.ForwardMsgID)
		params.Set("forwardChatId", message.ForwardChatID)
	}

	response, err := c.Do("/messages/sendFile", params, nil)
	if err != nil {
		return fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
		return fmt.Errorf("cannot unmarshal response: %s", err)
	}

	return nil
}

func (c *Client) SendVoiceMessage(message *Message) error {
	params := url.Values{
		"chatId":  {message.Chat.ID},
		"caption": {message.Text},
		"fileId":  {message.FileID},
	}

	if message.ReplyMsgID != "" {
		params.Set("replyMsgId", message.ReplyMsgID)
	}

	if message.ForwardMsgID != "" {
		params.Set("forwardMsgId", message.ForwardMsgID)
		params.Set("forwardChatId", message.ForwardChatID)
	}

	response, err := c.Do("/messages/sendVoice", params, nil)
	if err != nil {
		return fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
		return fmt.Errorf("cannot unmarshal response: %s", err)
	}

	return nil
}

func (c *Client) UploadFile(message *Message) error {
	params := url.Values{
		"chatId":  {message.Chat.ID},
		"caption": {message.Text},
	}

	response, err := c.Do("/messages/sendFile", params, message.File)
	if err != nil {
		return fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
		return fmt.Errorf("cannot unmarshal response: %s", err)
	}

	return nil
}

func (c *Client) UploadVoice(message *Message) error {
	params := url.Values{
		"chatId":  {message.Chat.ID},
		"caption": {message.Text},
	}

	response, err := c.Do("/messages/sendVoice", params, message.File)
	if err != nil {
		return fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(response, message); err != nil {
		return fmt.Errorf("cannot unmarshal response: %s", err)
	}

	return nil
}

func (c *Client) GetEvents(lastEventID int, pollTime int) ([]*Event, error) {
	params := url.Values{
		"lastEventId": {strconv.Itoa(lastEventID)},
		"pollTime":    {strconv.Itoa(pollTime)},
	}
	events := &eventsResponse{}

	response, err := c.Do("/events/get", params, nil)
	if err != nil {
		return events.Events, fmt.Errorf("error while making request: %s", err)
	}

	if err := json.Unmarshal(response, events); err != nil {
		return events.Events, fmt.Errorf("cannot parse events: %s", err)
	}

	return events.Events, nil
}

func (c *Client) PinMessage(message *Message) error {
	params := url.Values{
		"chatId": []string{message.Chat.ID},
		"msgId":  []string{message.ID},
	}
	_, err := c.Do("/chats/pinMessage", params, nil)
	if err != nil {
		return fmt.Errorf("error while pinning message: %s", err)
	}

	return nil
}

func (c *Client) UnpinMessage(message *Message) error {
	params := url.Values{
		"chatId": []string{message.Chat.ID},
		"msgId":  []string{message.ID},
	}
	_, err := c.Do("/chats/unpinMessage", params, nil)
	if err != nil {
		return fmt.Errorf("error while unpinning message: %s", err)
	}

	return nil
}

func (c *Client) SendAnswerCallbackQuery(answer *ButtonResponse) error {
	params := url.Values{
		"queryId":   []string{answer.QueryID},
		"text":      []string{answer.Text},
		"url":       []string{answer.URL},
		"showAlert": []string{strconv.FormatBool(answer.ShowAlert)},
	}

	_, err := c.Do("/messages/answerCallbackQuery", params, nil)
	if err != nil {
		return fmt.Errorf("error while making request: %s", err)
	}

	return nil
}

func NewClient(baseURL string, token string, logger *logrus.Logger) *Client {
	return &Client{
		token:   token,
		baseURL: baseURL,
		client:  http.DefaultClient,
		logger:  logger,
	}
}
