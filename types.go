package goicqbot

type BotInfo struct {
	// Id of the bot
	UserID string `json:"userId"`

	// Nickname of the bot
	Nick string `json:"nick"`

	// Name of the bot
	FirstName string `json:"firstName"`

	// Information about the box
	About string `json:"about"`

	// A slice of avatars
	Photo []string `json:"photo"`
}

type eventsResponse struct {
	OK     bool     `json:"ok"`
	Events []*Event `json:"events"`
}

type Contact struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type EventPayload struct {
	// Id of the message.
	// Presented in newMessage, editedMessage, deletedMessage, pinnedMessage, unpinnedMessage events.
	MsgID string `json:"msgId"`

	// Chat info.
	// Presented in all events.
	Chat ChatInfo `json:"chat"`

	// Author of the message
	// Presented in newMessage and editedMessage events.
	From Contact `json:"from"`

	// Text of the message.
	// Presented in newMessage, editedMessage and pinnedMessage events.
	Text string `json:"text"`

	// Parts of the message.
	// Presented only in newMessage event.
	Parts []Part `json:"parts"`

	// Timestamp of the event.
	Timestamp int `json:"timestamp"`

	LeftMembers []Contact `json:"leftMembers"`

	NewMembers []Contact `json:"newMembers"`

	AddedBy Contact `json:"addedBy"`

	RemovedBy Contact `json:"removedBy"`
}

type PartMessage struct {
	From      Contact `json:"from"`
	MsgID     string  `json:"msgId"`
	Text      string  `json:"text"`
	Timestamp int     `json:"timestamp"`
}

type PartPayload struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	UserID    string      `json:"userId"`
	FileID    string      `json:"fileId"`
	Caption   string      `json:"caption"`
	Type      string      `json:"type"`
	Message   PartMessage `json:"message"`
}

type Event struct {
	// Id of the event
	EventID int `json:"eventId"`

	// Type of the event: newMessage, editedMessage, deletedMessage, pinnedMessage, unpinnedMessage, newChatMembers
	Type string `json:"type"`

	// Payload of the event
	Payload EventPayload `json:"payload"`
}

type Part struct {
	// Type of the part
	Type string `json:"type"`

	// Payload of the part
	Payload PartPayload `json:"payload"`
}

type ChatInfo struct {
	// Id of the chat
	ChatID string `json:"chatId"`

	// Type of the chat: channel or group
	Type string `json:"type"`

	// Title of the chat
	Title string `json:"title"`
}

// Message represents a text message in ICQ
type Message struct {
	// Id of the message (for editing)
	MsgID string

	// Id of file to send
	FileID string

	// Text of the message or caption for file
	Text string

	// Chat where to send the message
	ChatID string

	// Id of replied message
	// You can't use it with ForwardMsgID or ForwardChatID
	ReplyMsgID string

	// Id of forwarded message
	// You can't use it with ReplyMsgID
	ForwardMsgID string

	// Id of a chat from which you forward the message
	// You can't use it with ReplyMsgID
	// You should use it with ForwardMsgID
	ForwardChatID string
}
