package botgolang

//go:generate easyjson -all types.go

type EventType string

type PartType string

const (
	NEW_MESSAGE       EventType = "newMessage"
	EDITED_MESSAGE    EventType = "editedMessage"
	DELETED_MESSAGE   EventType = "deletedMessage"
	PINNED_MESSAGE    EventType = "pinnedMessage"
	UNPINNED_MESSAGE  EventType = "unpinnedMessage"
	NEW_CHAT_MEMBERS  EventType = "newChatMembers"
	LEFT_CHAT_MEMBERS EventType = "leftChatMembers"
	CALLBACK_QUERY    EventType = "callbackQuery"

	STICKER PartType = "sticker"
	MENTION PartType = "mention"
	VOICE   PartType = "voice"
	FILE    PartType = "file"
	FORWARD PartType = "forward"
	REPLY   PartType = "reply"
)

type Response struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
}

type Thread struct {
	ThreadID string `json:"threadId"`
}

type UserState struct {
	Lastseen int `json:"lastseen"`
}

type Subscriber struct {
	SN        string    `json:"sn"`
	UserState UserState `json:"userState"`
}

type ThreadSubscribers struct {
	Cursor      string       `json:"cursor"`
	Subscribers []Subscriber `json:"subscribers"`
}

type Photo struct {
	URL string `json:"url"`
}

type BotInfo struct {
	User

	// Nickname of the bot
	Nick string `json:"nick"`

	// Name of the bot
	FirstName string `json:"firstName"`

	// Information about the box
	About string `json:"about"`

	// A slice of avatars
	Photo []Photo `json:"photo"`
}

type eventsResponse struct {
	OK     bool     `json:"ok"`
	Events []*Event `json:"events"`
}

type User struct {
	ID string `json:"userId"`
}

type ChatMember struct {
	User
	Creator bool `json:"creator"`
	Admin   bool `json:"admin"`
}

type UsersListResponse struct {
	List []User `json:"users"`
}

type MembersListResponse struct {
	// TODO: cursor
	List []ChatMember `json:"members"`
}

type AdminsListResponse struct {
	List []ChatMember `json:"admins"`
}

type Contact struct {
	User
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type BaseEventPayload struct {
	// Id of the message.
	// Presented in newMessage, editedMessage, deletedMessage, pinnedMessage, unpinnedMessage events.
	MsgID string `json:"msgId"`

	// Chat info.
	// Presented in all events.
	Chat Chat `json:"chat"`

	// Author of the message
	// Presented in newMessage and editedMessage events.
	From Contact `json:"from"`

	// Text of the message.
	// Presented in newMessage, editedMessage and pinnedMessage events.
	Text string `json:"text"`

	// Timestamp of the event.
	Timestamp int `json:"timestamp"`

	ParentMessage *ParentMessage `json:"parent_topic"`
}

type EventPayload struct {
	client *Client
	BaseEventPayload

	// Parts of the message.
	// Presented only in newMessage event.
	Parts []Part `json:"parts"`

	// Id of the query.
	// Presented only in callbackQuery event.
	QueryID string `json:"queryId"`

	// Callback message of the query (parent message for button).
	// Presented only in callbackQuery event.
	CallbackMsg BaseEventPayload `json:"message"`

	// CallbackData of the query (id of button).
	// Presented only in callbackQuery event.
	CallbackData string `json:"callbackData"`

	LeftMembers []Contact `json:"leftMembers"`

	NewMembers []Contact `json:"newMembers"`

	AddedBy Contact `json:"addedBy"`

	RemovedBy Contact `json:"removedBy"`
}

func (ep *EventPayload) Message() *Message {
	return message(ep.client, ep.BaseEventPayload)
}

func (ep *EventPayload) CallbackMessage() *Message {
	return message(ep.client, ep.CallbackMsg)
}

func message(client *Client, msg BaseEventPayload) *Message {
	msg.Chat.client = client
	return &Message{
		client:        client,
		ID:            msg.MsgID,
		Text:          msg.Text,
		Chat:          msg.Chat,
		Timestamp:     msg.Timestamp,
		ParentMessage: msg.ParentMessage,
	}
}

type PartMessage struct {
	From      Contact `json:"from"`
	MsgID     string  `json:"msgId"`
	Text      string  `json:"text"`
	Timestamp int     `json:"timestamp"`
}

type PartPayload struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	UserID      string      `json:"userId"`
	FileID      string      `json:"fileId"`
	Caption     string      `json:"caption"`
	Type        string      `json:"type"`
	PartMessage PartMessage `json:"message"`
	Message     PartMessage `json:"-"`
}

type Event struct {
	client *Client

	// Id of the event
	EventID int `json:"eventId"`

	// Type of the event: newMessage, editedMessage, deletedMessage, pinnedMessage, unpinnedMessage, newChatMembers
	Type EventType `json:"type"`

	// Payload of the event
	Payload EventPayload `json:"payload"`
}

type Part struct {
	// Type of the part
	Type PartType `json:"type"`

	// Payload of the part
	Payload PartPayload `json:"payload"`
}

func (ep *EventPayload) CallbackQuery() *ButtonResponse {
	return &ButtonResponse{
		client:       ep.client,
		QueryID:      ep.QueryID,
		CallbackData: ep.CallbackData,
	}
}
