package goicqbot

type BotInfo struct {
	UserID    string   `json:"userId"`
	Nick      string   `json:"nick"`
	FirstName string   `json:"firstName"`
	About     string   `json:"about"`
	Photo     []string `json:"photo"`
}

type EventsResponse struct {
	OK     bool     `json:"ok"`
	Events []*Event `json:"events"`
}

type Contact struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type EventPayload struct {
	MsgID     string   `json:"msgId"`
	Chat      ChatInfo `json:"chat"`
	From      Contact  `json:"from"`
	Text      string   `json:"text"`
	Parts     []Part   `json:"parts"`
	Timestamp int      `json:"timestamp"`
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
	EventID int          `json:"eventId"`
	Type    string       `json:"type"`
	Payload EventPayload `json:"payload"`
}

type Part struct {
	Type    string      `json:"type"`
	Payload PartPayload `json:"payload"`
}

type ChatInfo struct {
	ChatID string `json:"chatId"`
	Type   string `json:"type"`
	Title  string `json:"title"`
}

type MessageResponse struct {
	OK     bool   `json:"ok"`
	MsgID  string `json:"msgId"`
	FileID string `json:"fileId"`
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
