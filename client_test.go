package botgolang

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Do_OK(t *testing.T) {
	assert := assert.New(t)
	testServer := httptest.NewServer(&MockHandler{})
	defer func() { testServer.Close() }()

	client := Client{
		baseURL: testServer.URL,
		token:   "test_token",
		client:  http.DefaultClient,
		logger:  &logrus.Logger{},
	}

	bytes, err := client.Do("/", url.Values{}, nil)

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

	bytes, err := client.Do("/", url.Values{}, nil)

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
			Type:    NEW_MESSAGE,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					MsgID: "57883346846815030",
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "channel",
						Title: "The best channel",
					},
					From: Contact{
						User:      User{"1234567890"},
						FirstName: "Name",
						LastName:  "SurName",
					},
					Text:      "Hello!",
					Timestamp: 1546290000,
				},
				Parts: []Part{
					{
						Type: STICKER,
						Payload: PartPayload{
							FileID: "2IWuJzaNWCJZxJWCvZhDYuJ5XDsr7hU",
						},
					},
					{
						Type: MENTION,
						Payload: PartPayload{
							FirstName: "Name",
							LastName:  "SurName",
							UserID:    "1234567890",
						},
					},
					{
						Type: VOICE,
						Payload: PartPayload{
							FileID: "IdjUEXuGdNhLKUfD5rvkE03IOax54cD",
						},
					},
					{
						Type: FILE,
						Payload: PartPayload{
							FileID:  "ZhSnMuaOmF7FRez2jGWuQs5zGZwlLa0",
							Caption: "Last weekend trip",
							Type:    "image",
						},
					},
					{
						Type: FORWARD,
						Payload: PartPayload{
							PartMessage: PartMessage{
								MsgID: "12354",
								Text:  "test1",
							},
						},
					},
					{
						Type: REPLY,
						Payload: PartPayload{
							PartMessage: PartMessage{
								MsgID: "12354",
								Text:  "test",
							},
						},
					},
				},
			},
		},
		{
			EventID: 2,
			Type:    EDITED_MESSAGE,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					MsgID: "57883346846815030",
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "channel",
						Title: "The best channel",
					},
					From: Contact{
						User:      User{"1234567890"},
						FirstName: "Name",
						LastName:  "SurName",
					},
					Text:      "Hello!",
					Timestamp: 1546290000,
				},
			},
		},
		{
			EventID: 3,
			Type:    DELETED_MESSAGE,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					MsgID: "57883346846815030",
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "channel",
						Title: "The best channel",
					},
					Timestamp: 1546290000,
				},
			},
		},
		{
			EventID: 4,
			Type:    PINNED_MESSAGE,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					MsgID: "6720509406122810000",
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "group",
						Title: "The best group",
					},
					From: Contact{
						User:      User{"9876543210"},
						FirstName: "Name",
						LastName:  "SurName",
					},
					Text:      "Some important information!",
					Timestamp: 1564740530,
				},
			},
		},
		{
			EventID: 5,
			Type:    UNPINNED_MESSAGE,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					MsgID: "6720509406122810000",
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "group",
						Title: "The best group",
					},
					Timestamp: 1564740530,
				},
			},
		},
		{
			EventID: 6,
			Type:    NEW_CHAT_MEMBERS,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "group",
						Title: "The best group",
					},
				},
				NewMembers: []Contact{
					{
						User:      User{"1234567890"},
						FirstName: "Name",
						LastName:  "SurName",
					},
				},
				AddedBy: Contact{
					User:      User{"9876543210"},
					FirstName: "Name",
					LastName:  "SurName",
				},
			},
		},
		{
			EventID: 7,
			Type:    LEFT_CHAT_MEMBERS,
			Payload: EventPayload{
				BaseEventPayload: BaseEventPayload{
					Chat: Chat{
						ID:    "681869378@chat.agent",
						Type:  "group",
						Title: "The best group",
					},
				},
				LeftMembers: []Contact{
					{
						User:      User{"1234567890"},
						FirstName: "Name",
						LastName:  "SurName",
					},
				},
				RemovedBy: Contact{
					User:      User{"9876543210"},
					FirstName: "Name",
					LastName:  "SurName",
				},
			},
		},
		{
			EventID: 8,
			Type:    CALLBACK_QUERY,
			Payload: EventPayload{
				CallbackData: "echo",
				CallbackMsg: BaseEventPayload{
					MsgID: "6720509406122810000",
					Chat: Chat{
						ID:   "1234567890",
						Type: "private",
					},
					From: Contact{
						User:      User{"bot_id"},
						FirstName: "bot_name",
					},
					Text:      "Some important information!",
					Timestamp: 1564740530,
				},
				BaseEventPayload: BaseEventPayload{
					From: Contact{
						User:      User{"1234567890"},
						FirstName: "Name",
					},
				},
				QueryID: "SVR:123456",
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

func TestClient_GetInfo_OK(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	testServer := httptest.NewServer(&MockHandler{})
	defer func() { testServer.Close() }()

	client := Client{
		baseURL: testServer.URL,
		token:   "test_token",
		client:  http.DefaultClient,
		logger:  &logrus.Logger{},
	}

	info, err := client.GetChatInfo("id_1234")
	require.NoError(err)
	assert.NotEmpty(info.ID)
}

func TestClient_GetInfo_Error(t *testing.T) {
	require := require.New(t)

	require.NoError(nil)
}
