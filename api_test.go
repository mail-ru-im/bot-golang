package goicqbot

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_Do_OK(t *testing.T) {
	assert := assert.New(t)
	testServer := httptest.NewServer(&MockHandler{})
	defer func() { testServer.Close() }()

	client := Client{
		baseURL: testServer.URL,
		token:   "test",
		client:  http.DefaultClient,
		logger:  &logrus.Logger{},
	}

	bytes, err := client.Do("/", url.Values{})

	assert.NoError(err)
	assert.JSONEq(`{"ok":true}`, string(bytes))
}

func TestClient_Do_Error(t *testing.T) {
	assert := assert.New(t)
	testServer := httptest.NewServer(&MockHandler{})
	defer func() { testServer.Close() }()

	client := Client{
		baseURL: testServer.URL,
		token:   "",
		client:  http.DefaultClient,
		logger:  &logrus.Logger{},
	}

	expected := `{"ok":false, "description":"Missing required parameter 'token'"}`

	bytes, err := client.Do("/", url.Values{})

	assert.EqualError(err, "error status from API: Missing required parameter 'token'")
	assert.JSONEq(expected, string(bytes))
}

func TestClient_GetEvents_OK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	testServer := httptest.NewServer(&MockHandler{})
	defer func() { testServer.Close() }()

	expected := []*Event{
		{
			EventID: 1,
			Type:    "newMessage",
			Payload: EventPayload{
				MsgID: "57883346846815030",
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "channel",
					Title:  "The best channel",
				},
				From: Contact{
					UserID:    "1234567890",
					FirstName: "Name",
					LastName:  "SurName",
				},
				Text: "Hello!",
				Parts: []Part{
					{
						Type: "sticker",
						Payload: PartPayload{
							FileID: "2IWuJzaNWCJZxJWCvZhDYuJ5XDsr7hU",
						},
					},
					{
						Type: "mention",
						Payload: PartPayload{
							FirstName: "Name",
							LastName:  "SurName",
							UserID:    "1234567890",
						},
					},
					{
						Type: "voice",
						Payload: PartPayload{
							FileID: "IdjUEXuGdNhLKUfD5rvkE03IOax54cD",
						},
					},
					{
						Type: "file",
						Payload: PartPayload{
							FileID:  "ZhSnMuaOmF7FRez2jGWuQs5zGZwlLa0",
							Caption: "Last weekend trip",
							Type:    "image",
						},
					},
					{
						Type: "forward",
						Payload: PartPayload{
							Message: "Some message to forward",
						},
					},
					{
						Type: "reply",
						Payload: PartPayload{
							Message: "Some message to reply",
						},
					},
				},
				Timestamp: 1546290000,
			},
		},
		{
			EventID: 2,
			Type:    "editedMessage",
			Payload: EventPayload{
				MsgID: "57883346846815030",
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "channel",
					Title:  "The best channel",
				},
				From: Contact{
					UserID:    "1234567890",
					FirstName: "Name",
					LastName:  "SurName",
				},
				Text:      "Hello!",
				Timestamp: 1546290000,
			},
		},
		{
			EventID: 3,
			Type:    "deletedMessage",
			Payload: EventPayload{
				MsgID: "57883346846815030",
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "channel",
					Title:  "The best channel",
				},
				Timestamp: 1546290000,
			},
		},
		{
			EventID: 4,
			Type:    "pinnedMessage",
			Payload: EventPayload{
				MsgID: "6720509406122810000",
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "group",
					Title:  "The best group",
				},
				From: Contact{
					UserID:    "9876543210",
					FirstName: "Name",
					LastName:  "SurName",
				},
				Text:      "Some important information!",
				Timestamp: 1564740530,
			},
		},
		{
			EventID: 5,
			Type:    "unpinnedMessage",
			Payload: EventPayload{
				MsgID: "6720509406122810000",
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "group",
					Title:  "The best group",
				},
				Timestamp: 1564740530,
			},
		},
		{
			EventID: 6,
			Type:    "newChatMembers",
			Payload: EventPayload{
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "group",
					Title:  "The best group",
				},
			},
		},
		{
			EventID: 7,
			Type:    "leftChatMembers",
			Payload: EventPayload{
				Chat: ChatInfo{
					ChatID: "681869378@chat.agent",
					Type:   "group",
					Title:  "The best group",
				},
			},
		},
	}

	client := Client{
		baseURL: testServer.URL,
		token:   "test_token",
		client:  http.DefaultClient,
		logger:  &logrus.Logger{},
	}

	events, err := client.GetEvents(0, 0)

	require.NoError(err)
	assert.Equal(events, expected)
}
