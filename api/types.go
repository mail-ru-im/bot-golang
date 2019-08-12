package api

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

type PartPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserID    string `json:"userId"`
	FileID    string `json:"fileId"`
	Caption   string `json:"caption"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type Event struct {
	EventID int     `json:"eventId"`
	Type    string  `json:"type"`
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
