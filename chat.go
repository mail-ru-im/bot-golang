package goicqbot

//go:generate easyjson -all chat.go

type Chat struct {
	client *Client
	// Id of the chat
	ID string `json:"chatId"`

	// Type of the chat: channel or group
	Type string `json:"type"`

	// Title of the chat
	Title string `json:"title"`

	// Is this chat public?
	Public bool `json:"public"`

	Group string `json:"group"`

	// You can send this link to all your friends
	InviteLink string `json:"inviteLink"`

	Admins []Contact `json:"admins"`
}
