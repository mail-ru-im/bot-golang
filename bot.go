package botgolang

/*
💥 botgolang is zero-configuration library with convenient interface.
Crafted with love in @mail for your awesome bots.
*/

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	defaultAPIURL = "https://api.icq.net/bot/v1"
	defaultDebug  = false
)

// Bot is the main structure for interaction with API.
// All fields are private, you can configure bot using config arguments in NewBot func.
type Bot struct {
	client  *Client
	updater *Updater
	logger  *logrus.Logger
	Info    *BotInfo
}

// GetInfo returns information about bot:
// id, name, about, avatar
//
// GetInfo uses context.Background internally; to specify the context, use
// GetInfoWithContext.
func (b *Bot) GetInfo() (*BotInfo, error) {
	return b.GetInfoWithContext(context.Background())
}

// GetInfoWithContext returns information about bot:
// id, name, about, avatar
func (b *Bot) GetInfoWithContext(ctx context.Context) (*BotInfo, error) {
	return b.client.GetInfoWithContext(ctx)
}

// GetChatInfo returns information about chat:
// id, type, title, public, group, inviteLink, admins
//
// GetChatInfo uses context.Background internally; to specify the context, use
// GetChatInfoWithContext.
func (b *Bot) GetChatInfo(chatID string) (*Chat, error) {
	return b.GetChatInfoWithContext(context.Background(), chatID)
}

// GetChatInfoWithContext returns information about chat:
// id, type, title, public, group, inviteLink, admins
func (b *Bot) GetChatInfoWithContext(ctx context.Context, chatID string) (*Chat, error) {
	return b.client.GetChatInfoWithContext(ctx, chatID)
}

// SendChatActions sends an actions like "typing, looking"
//
// SendChatActions uses context.Background internally; to specify the context, use
// SendChatActionsWithContext.
func (b *Bot) SendChatActions(chatID string, actions ...ChatAction) error {
	return b.SendChatActionsWithContext(context.Background(), chatID, actions...)
}

// SendChatActionsWithContext sends an actions like "typing, looking"
func (b *Bot) SendChatActionsWithContext(ctx context.Context, chatID string, actions ...ChatAction) error {
	return b.client.SendChatActionsWithContext(ctx, chatID, actions...)
}

// GetChatAdmins returns chat admins list with fields:
// userID, creator flag
//
// GetChatAdmins uses context.Background internally; to specify the context, use
// GetChatAdminsWithContext.
func (b *Bot) GetChatAdmins(chatID string) ([]ChatMember, error) {
	return b.GetChatAdminsWithContext(context.Background(), chatID)
}

// GetChatAdminsWithContext returns chat admins list with fields:
// userID, creator flag
func (b *Bot) GetChatAdminsWithContext(ctx context.Context, chatID string) ([]ChatMember, error) {
	return b.client.GetChatAdminsWithContext(ctx, chatID)
}

// GetChatMem returns chat members list with fields:
// userID, creator flag, admin flag
//
// GetChatMembers uses context.Background internally; to specify the context, use
// GetChatMembersWithContext.
func (b *Bot) GetChatMembers(chatID string) ([]ChatMember, error) {
	return b.GetChatMembersWithContext(context.Background(), chatID)
}

// GetChatMembersWithContext returns chat members list with fields:
// userID, creator flag, admin flag
func (b *Bot) GetChatMembersWithContext(ctx context.Context, chatID string) ([]ChatMember, error) {
	return b.client.GetChatMembersWithContext(ctx, chatID)
}

// GetChatBlockedUsers returns chat blocked users list:
// userID
//
// GetChatBlockedUsers uses context.Background internally; to specify the context, use
// GetChatBlockedUsersWithContext.
func (b *Bot) GetChatBlockedUsers(chatID string) ([]User, error) {
	return b.GetChatBlockedUsersWithContext(context.Background(), chatID)
}

// GetChatBlockedUsersWithContext returns chat blocked users list:
// userID
func (b *Bot) GetChatBlockedUsersWithContext(ctx context.Context, chatID string) ([]User, error) {
	return b.client.GetChatBlockedUsersWithContext(ctx, chatID)
}

// GetChatPendingUsers returns chat join pending users list:
// userID
//
// GetChatPendingUsers uses context.Background internally; to specify the context, use
// GetChatPendingUsersWithContext.
func (b *Bot) GetChatPendingUsers(chatID string) ([]User, error) {
	return b.GetChatPendingUsersWithContext(context.Background(), chatID)
}

// GetChatPendingUsersWithContext returns chat join pending users list:
// userID
func (b *Bot) GetChatPendingUsersWithContext(ctx context.Context, chatID string) ([]User, error) {
	return b.client.GetChatPendingUsersWithContext(ctx, chatID)
}

// BlockChatUser blocks user and removes him from chat.
// If deleteLastMessages is true, the messages written recently will be deleted
//
// BlockChatUser uses context.Background internally; to specify the context, use
// BlockChatUserWithContext.
func (b *Bot) BlockChatUser(chatID, userID string, deleteLastMessages bool) error {
	return b.BlockChatUserWithContext(context.Background(), chatID, userID, deleteLastMessages)
}

// BlockChatUserWithContext blocks user and removes him from chat.
// If deleteLastMessages is true, the messages written recently will be deleted
func (b *Bot) BlockChatUserWithContext(ctx context.Context, chatID, userID string, deleteLastMessages bool) error {
	return b.client.BlockChatUserWithContext(ctx, chatID, userID, deleteLastMessages)
}

// UnblockChatUser unblocks user in chat
//
// UnblockChatUser uses context.Background internally; to specify the context, use
// UnblockChatUserWithContext.
func (b *Bot) UnblockChatUser(chatID, userID string) error {
	return b.UnblockChatUserWithContext(context.Background(), chatID, userID)
}

// UnblockChatUserWithContext unblocks user in chat
func (b *Bot) UnblockChatUserWithContext(ctx context.Context, chatID, userID string) error {
	return b.client.UnblockChatUserWithContext(ctx, chatID, userID)
}

// ResolveChatJoinRequests sends a decision to accept/decline user join to chat
//
// ResolveChatJoinRequests uses context.Background internally; to specify the context, use
// ResolveChatJoinRequestsWithContext.
func (b *Bot) ResolveChatJoinRequests(chatID, userID string, accept, everyone bool) error {
	return b.ResolveChatJoinRequestsWithContext(context.Background(), chatID, userID, accept, everyone)
}

// ResolveChatJoinRequestsWithContext sends a decision to accept/decline user join to chat
func (b *Bot) ResolveChatJoinRequestsWithContext(ctx context.Context, chatID, userID string, accept, everyone bool) error {
	return b.client.ResolveChatPendingWithContext(ctx, chatID, userID, accept, everyone)
}

// SetChatTitle changes chat title
//
// SetChatTitle uses context.Background internally; to specify the context, use
// SetChatTitleWithContext.
func (b *Bot) SetChatTitle(chatID, title string) error {
	return b.SetChatTitleWithContext(context.Background(), chatID, title)
}

// SetChatTitleWithContext changes chat title
func (b *Bot) SetChatTitleWithContext(ctx context.Context, chatID, title string) error {
	return b.client.SetChatTitleWithContext(ctx, chatID, title)
}

// SetChatAbout changes chat about
//
// SetChatAbout uses context.Background internally; to specify the context, use
// SetChatAboutWithContext.
func (b *Bot) SetChatAbout(chatID, about string) error {
	return b.SetChatAboutWithContext(context.Background(), chatID, about)
}

// SetChatAboutWithContext changes chat about
func (b *Bot) SetChatAboutWithContext(ctx context.Context, chatID, about string) error {
	return b.client.SetChatAboutWithContext(ctx, chatID, about)
}

// SetChatRules changes chat rules
//
// SetChatRules uses context.Background internally; to specify the context, use
// SetChatRulesWithContext.
func (b *Bot) SetChatRules(chatID, rules string) error {
	return b.SetChatRulesWithContext(context.Background(), chatID, rules)
}

// SetChatRulesWithContext changes chat rules
func (b *Bot) SetChatRulesWithContext(ctx context.Context, chatID, rules string) error {
	return b.client.SetChatRulesWithContext(ctx, chatID, rules)
}

// GetFileInfo returns information about file:
// id, type, size, filename, url
//
// GetFileInfo uses context.Background internally; to specify the context, use
// GetFileInfoWithContext.
func (b *Bot) GetFileInfo(fileID string) (*File, error) {
	return b.GetFileInfoWithContext(context.Background(), fileID)
}

// GetFileInfoWithContext returns information about file:
// id, type, size, filename, url
func (b *Bot) GetFileInfoWithContext(ctx context.Context, fileID string) (*File, error) {
	return b.client.GetFileInfoWithContext(ctx, fileID)
}

// NewMessage returns new message
func (b *Bot) NewMessage(chatID string) *Message {
	return &Message{
		client: b.client,
		Chat:   Chat{ID: chatID},
	}
}

// NewTextMessage returns new text message
func (b *Bot) NewTextMessage(chatID, text string) *Message {
	return &Message{
		client:      b.client,
		Chat:        Chat{ID: chatID},
		Text:        text,
		ContentType: Text,
	}
}

// NewInlineKeyboardMessage returns new text message with inline keyboard
func (b *Bot) NewInlineKeyboardMessage(chatID, text string, keyboard Keyboard) *Message {
	return &Message{
		client:         b.client,
		Chat:           Chat{ID: chatID},
		Text:           text,
		ContentType:    Text,
		InlineKeyboard: &keyboard,
	}
}

// NewFileMessage returns new file message
func (b *Bot) NewFileMessage(chatID string, file *os.File) *Message {
	return &Message{
		client:      b.client,
		Chat:        Chat{ID: chatID},
		File:        file,
		ContentType: OtherFile,
	}
}

// NewFileMessageByFileID returns new message with previously uploaded file id
func (b *Bot) NewFileMessageByFileID(chatID, fileID string) *Message {
	return &Message{
		client:      b.client,
		Chat:        Chat{ID: chatID},
		FileID:      fileID,
		ContentType: OtherFile,
	}
}

// NewVoiceMessage returns new voice message
func (b *Bot) NewVoiceMessage(chatID string, file *os.File) *Message {
	return &Message{
		client:      b.client,
		Chat:        Chat{ID: chatID},
		File:        file,
		ContentType: Voice,
	}
}

// NewVoiceMessageByFileID returns new message with previously uploaded voice file id
func (b *Bot) NewVoiceMessageByFileID(chatID, fileID string) *Message {
	return &Message{
		client:      b.client,
		Chat:        Chat{ID: chatID},
		FileID:      fileID,
		ContentType: Voice,
	}
}

// NewMessageFromPart returns new message based on part message
func (b *Bot) NewMessageFromPart(message PartMessage) *Message {
	return &Message{
		client:    b.client,
		ID:        message.MsgID,
		Chat:      Chat{ID: message.From.User.ID, Title: message.From.FirstName},
		Text:      message.Text,
		Timestamp: message.Timestamp,
	}
}

// NewButtonResponse returns new ButtonResponse
func (b *Bot) NewButtonResponse(queryID, url, text string, showAlert bool) *ButtonResponse {
	return &ButtonResponse{
		client:    b.client,
		QueryID:   queryID,
		Text:      text,
		URL:       url,
		ShowAlert: showAlert,
	}
}

func (b *Bot) NewChat(id string) *Chat {
	return &Chat{
		client: b.client,
		ID:     id,
	}
}

// SendMessage sends a message, passed as an argument.
// This method fills the argument with ID of sent message and returns an error if any.
//
// SendMessage uses context.Background internally; to specify the context, use
// SendMessageWithContext.
func (b *Bot) SendMessage(message *Message) error {
	return b.SendMessageWithContext(context.Background(), message)
}

// SendMessageWithContext sends a message, passed as an argument.
// This method fills the argument with ID of sent message and returns an error if any.
func (b *Bot) SendMessageWithContext(ctx context.Context, message *Message) error {
	message.client = b.client
	return message.SendWithContext(ctx)
}

// EditMessage edit a message passed as an argument.
//
// EditMessage uses context.Background internally; to specify the context, use
// EditMessageWithContext.
func (b *Bot) EditMessage(message *Message) error {
	return b.EditMessageWithContext(context.Background(), message)
}

// EditMessageWithContext edit a message passed as an argument.
func (b *Bot) EditMessageWithContext(ctx context.Context, message *Message) error {
	return b.client.EditMessageWithContext(ctx, message)
}

// GetUpdatesChannel returns a channel, which will be filled with events.
// You can pass cancellable context there and stop receiving events.
// The channel will be closed after context cancellation.
func (b *Bot) GetUpdatesChannel(ctx context.Context) <-chan Event {
	updates := make(chan Event)

	go b.updater.RunUpdatesCheck(ctx, updates)

	return updates
}

// NewBot returns new bot object.
// All communications with bot API must go through Bot struct.
// In general you don't need to configure this bot, therefore all options are optional arguments.
func NewBot(token string, opts ...BotOption) (*Bot, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	apiURL := defaultAPIURL
	debug := defaultDebug
	client := *http.DefaultClient
	for _, option := range opts {
		switch option.Type() {
		case "api_url":
			apiURL = option.Value().(string)
		case "debug":
			debug = option.Value().(bool)
		case "http_client":
			client = option.Value().(http.Client)
		}
	}

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	tgClient := NewCustomClient(&client, apiURL, token, logger)
	updater := NewUpdater(tgClient, 0, logger)

	info, err := tgClient.GetInfo()
	if err != nil {
		return nil, fmt.Errorf("cannot get info about bot: %s", err)
	}

	return &Bot{
		client:  tgClient,
		updater: updater,
		logger:  logger,
		Info:    info,
	}, nil
}
