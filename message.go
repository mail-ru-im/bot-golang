package botgolang

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

//go:generate easyjson -all message.go

type MessageContentType uint8

const (
	Unknown MessageContentType = iota
	Text
	OtherFile
	Voice
)

// Message represents a text message
type Message struct {
	client      *Client
	ContentType MessageContentType

	// Id of the message (for editing)
	ID string `json:"msgId"`

	// File contains file attachment of the message
	File *os.File `json:"-"`

	// Id of file to send
	FileID string `json:"fileId"`

	// Text of the message or caption for file
	Text string `json:"text"`

	// Chat where to send the message
	Chat Chat `json:"chat"`

	// Id of replied message
	// You can't use it with ForwardMsgID or ForwardChatID
	ReplyMsgID string `json:"replyMsgId"`

	// Id of forwarded message
	// You can't use it with ReplyMsgID
	ForwardMsgID string `json:"forwardMsgId"`

	// Id of a chat from which you forward the message
	// You can't use it with ReplyMsgID
	// You should use it with ForwardMsgID
	ForwardChatID string `json:"forwardChatId"`

	Timestamp int `json:"timestamp"`

	// The markup for the inline keyboard
	InlineKeyboard *Keyboard `json:"inlineKeyboardMarkup"`

	// The parse mode (HTML/MarkdownV2)
	ParseMode ParseMode `json:"parseMode"`
}

func (m *Message) AttachNewFile(file *os.File) {
	m.File = file
	m.ContentType = OtherFile
}

func (m *Message) AttachExistingFile(fileID string) {
	m.FileID = fileID
	m.ContentType = OtherFile
}

func (m *Message) AttachNewVoice(file *os.File) {
	m.File = file
	m.ContentType = Voice
}

func (m *Message) AttachExistingVoice(fileID string) {
	m.FileID = fileID
	m.ContentType = Voice
}

// ParseMode represent a type of text formatting
type ParseMode string

const (
	ParseModeHTML       ParseMode = "HTML"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
)

// AppendParseMode append a type of text formatting for current message
func (m *Message) AppendParseMode(mode ParseMode) {
	m.ParseMode = mode
}

// AttachInlineKeyboard adds a keyboard to the message.
// Note - at least one row should be in the keyboard
// and there should be no empty rows
func (m *Message) AttachInlineKeyboard(keyboard Keyboard) {
	m.InlineKeyboard = &keyboard
}

// Send method sends your message.
// Make sure you have Text or FileID in your message.
//
// Send uses context.Background internally; to specify the context, use
// SendWithContext.
func (m *Message) Send() error {
	return m.SendWithContext(context.Background())
}

// SendWithContext method sends your message.
// Make sure you have Text or FileID in your message.
func (m *Message) SendWithContext(ctx context.Context) error {
	if m.client == nil {
		return fmt.Errorf("client is not inited, create message with constructor NewMessage, NewTextMessage, etc")
	}

	if m.Chat.ID == "" {
		return fmt.Errorf("message should have chat id")
	}

	switch m.ContentType {
	case Voice:
		if m.FileID != "" {
			return m.client.SendVoiceMessageWithContext(ctx, m)
		}

		if m.File != nil {
			return m.client.UploadVoiceWithContext(ctx, m)
		}
	case OtherFile:
		if m.FileID != "" {
			return m.client.SendFileMessageWithContext(ctx, m)
		}

		if m.File != nil {
			return m.client.UploadFileWithContext(ctx, m)
		}
	case Text:
		return m.client.SendTextMessageWithContext(ctx, m)
	case Unknown:
		// need to autodetect
		if m.FileID != "" {
			// voice message's fileID always starts with 'I'
			if m.FileID[0] == voiceMessageLeadingRune {
				return m.client.SendVoiceMessageWithContext(ctx, m)
			}
			return m.client.SendFileMessageWithContext(ctx, m)
		}

		if m.File != nil {
			if voiceMessageSupportedExtensions[filepath.Ext(m.File.Name())] {
				return m.client.UploadVoiceWithContext(ctx, m)
			}
			return m.client.UploadFileWithContext(ctx, m)
		}

		if m.Text != "" {
			return m.client.SendTextMessageWithContext(ctx, m)
		}
	}

	return fmt.Errorf("cannot send message or file without data")
}

// Edit method edits your message.
// Make sure you have ID in your message.
//
// Edit uses context.Background internally; to specify the context, use
// EditWithContext.
func (m *Message) Edit() error {
	return m.EditWithContext(context.Background())
}

// EditWithContext method edits your message.
// Make sure you have ID in your message.
func (m *Message) EditWithContext(ctx context.Context) error {
	if m.ID == "" {
		return fmt.Errorf("cannot edit message without id")
	}
	return m.client.EditMessageWithContext(ctx, m)
}

// Delete method deletes your message.
// Make sure you have ID in your message.
//
// Delete uses context.Background internally; to specify the context, use
// DeleteWithContext.
func (m *Message) Delete() error {
	return m.DeleteWithContext(context.Background())
}

// DeleteWithContext method deletes your message.
// Make sure you have ID in your message.
func (m *Message) DeleteWithContext(ctx context.Context) error {
	if m.ID == "" {
		return fmt.Errorf("cannot delete message without id")
	}

	return m.client.DeleteMessageWithContext(ctx, m)
}

// Reply method replies to the message.
// Make sure you have ID in the message.
//
// Reply uses context.Background internally; to specify the context, use
// ReplyWithContext.
func (m *Message) Reply(text string) error {
	return m.ReplyWithContext(context.Background(), text)
}

// ReplyWithContext method replies to the message.
// Make sure you have ID in the message.
func (m *Message) ReplyWithContext(ctx context.Context, text string) error {
	if m.ID == "" {
		return fmt.Errorf("cannot reply to message without id")
	}

	m.ReplyMsgID = m.ID
	m.Text = text

	return m.client.SendTextMessageWithContext(ctx, m)
}

// Forward method forwards your message to chat.
// Make sure you have ID in your message.
//
// Forward uses context.Background internally; to specify the context, use
// ForwardWithContext.
func (m *Message) Forward(chatID string) error {
	return m.ForwardWithContext(context.Background(), chatID)
}

// ForwardWithContext method forwards your message to chat.
// Make sure you have ID in your message.
func (m *Message) ForwardWithContext(ctx context.Context, chatID string) error {
	if m.ID == "" {
		return fmt.Errorf("cannot forward message without id")
	}

	m.ForwardChatID = m.Chat.ID
	m.ForwardMsgID = m.ID
	m.Chat.ID = chatID

	return m.client.SendTextMessageWithContext(ctx, m)
}

// Pin message in chat
// Make sure you are admin in this chat
//
// Pin uses context.Background internally; to specify the context, use
// PinWithContext.
func (m *Message) Pin() error {
	return m.PinWithContext(context.Background())
}

// PinWithContext message in chat
// Make sure you are admin in this chat
func (m *Message) PinWithContext(ctx context.Context) error {
	if m.ID == "" {
		return fmt.Errorf("cannot pin message without id")
	}

	return m.client.PinMessageWithContext(ctx, m)
}

// Unpin message in chat
// Make sure you are admin in this chat
//
// Unpin uses context.Background internally; to specify the context, use
// UnpinWithContext.
func (m *Message) Unpin() error {
	return m.UnpinWithContext(context.Background())
}

// Unpin message in chat
// Make sure you are admin in this chat
func (m *Message) UnpinWithContext(ctx context.Context) error {
	if m.ID == "" {
		return fmt.Errorf("cannot unpin message without id")
	}

	return m.client.UnpinMessageWithContext(ctx, m)
}
