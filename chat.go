package botgolang

import "context"

//go:generate easyjson -all chat.go

type ChatAction = string

const (
	TypingAction  ChatAction = "typing"
	LookingAction ChatAction = "looking"
)

type ChatType = string

const (
	Private ChatType = "private"
	Group   ChatType = "group"
	Channel ChatType = "channel"
)

type Chat struct {
	client *Client
	// Id of the chat
	ID string `json:"chatId"`

	// Type of the chat: channel, group or private
	Type ChatType `json:"type"`

	// First name of the user
	FirstName string `json:"firstName"`

	// Last name of the user
	LastName string `json:"lastName"`

	// Nick of the user
	Nick string `json:"nick"`

	// User about or group/channel description
	About string `json:"about"`

	// Rules of the group/channel
	Rules string `json:"rules"`

	// Title of the chat
	Title string `json:"title"`

	// Flag that indicates that requested chat is the bot
	IsBot bool `json:"isBot"`

	// Is this chat public?
	Public bool `json:"public"`

	// Is this chat has join moderation?
	JoinModeration bool `json:"joinModeration"`

	// You can send this link to all your friends
	InviteLink string `json:"inviteLink"`
}

func (c *Chat) resolveID() string {
	switch c.Type {
	case Private:
		return c.Nick
	default:
		return c.ID
	}
}

// Send bot actions to the chat
//
// You can call this method every time you change the current actions,
// or every 10 seconds if the actions have not changed. After sending a
// request without active action, you should not re-notify of their absence.
//
// SendActions uses context.Background internally; to specify the context, use
// SendActionsWithContext.
func (c *Chat) SendActions(actions ...ChatAction) error {
	return c.SendActionsWithContext(context.Background(), actions...)
}

// Send bot actions to the chat
//
// You can call this method every time you change the current actions,
// or every 10 seconds if the actions have not changed. After sending a
// request without active action, you should not re-notify of their absence.
func (c *Chat) SendActionsWithContext(ctx context.Context, actions ...ChatAction) error {
	return c.client.SendChatActionsWithContext(ctx, c.resolveID(), actions...)
}

// Get chat administrators list
//
// GetAdmins uses context.Background internally; to specify the context, use
// GetAdminsWithContext.
func (c *Chat) GetAdmins() ([]ChatMember, error) {
	return c.GetAdminsWithContext(context.Background())
}

// Get chat administrators list
func (c *Chat) GetAdminsWithContext(ctx context.Context) ([]ChatMember, error) {
	return c.client.GetChatAdminsWithContext(ctx, c.ID)
}

// Get chat members list
//
// GetMembers uses context.Background internally; to specify the context, use
// GetMembersWithContext.
func (c *Chat) GetMembers() ([]ChatMember, error) {
	return c.GetMembersWithContext(context.Background())
}

// Get chat members list
func (c *Chat) GetMembersWithContext(ctx context.Context) ([]ChatMember, error) {
	return c.client.GetChatMembersWithContext(ctx, c.ID)
}

// Get chat blocked users list
//
// GetBlockedUsers uses context.Background internally; to specify the context, use
// GetBlockedUsersWithContext.
func (c *Chat) GetBlockedUsers() ([]User, error) {
	return c.GetBlockedUsersWithContext(context.Background())
}

// Get chat blocked users list
func (c *Chat) GetBlockedUsersWithContext(ctx context.Context) ([]User, error) {
	return c.client.GetChatBlockedUsersWithContext(ctx, c.ID)
}

// Get chat join pending users list
//
// GetPendingUsers uses context.Background internally; to specify the context, use
// GetPendingUsersWithContext.
func (c *Chat) GetPendingUsers() ([]User, error) {
	return c.GetPendingUsersWithContext(context.Background())
}

// Get chat join pending users list
func (c *Chat) GetPendingUsersWithContext(ctx context.Context) ([]User, error) {
	return c.client.GetChatPendingUsersWithContext(ctx, c.ID)
}

// Block user and remove him from chat.
// If deleteLastMessages is true, the messages written recently will be deleted
//
// BlockUser uses context.Background internally; to specify the context, use
// BlockUserWithContext.
func (c *Chat) BlockUser(userID string, deleteLastMessages bool) error {
	return c.BlockUserWithContext(context.Background(), userID, deleteLastMessages)
}

// Block user and remove him from chat.
// If deleteLastMessages is true, the messages written recently will be deleted
func (c *Chat) BlockUserWithContext(ctx context.Context, userID string, deleteLastMessages bool) error {
	return c.client.BlockChatUserWithContext(ctx, c.ID, userID, deleteLastMessages)
}

// Unblock user in chat (but not add him back)
//
// UnblockUser uses context.Background internally; to specify the context, use
// UnblockUserWithContext.
func (c *Chat) UnblockUser(userID string) error {
	return c.UnblockUserWithContext(context.Background(), userID)
}

// Unblock user in chat (but not add him back)
func (c *Chat) UnblockUserWithContext(ctx context.Context, userID string) error {
	return c.client.UnblockChatUserWithContext(ctx, c.ID, userID)
}

// ResolveJoinRequest resolve specific user chat join request
//
// ResolveJoinRequest uses context.Background internally; to specify the context, use
// ResolveJoinRequestWithContext.
func (c *Chat) ResolveJoinRequest(userID string, accept bool) error {
	return c.ResolveJoinRequestWithContext(context.Background(), userID, accept)
}

// ResolveJoinRequest resolve specific user chat join request
func (c *Chat) ResolveJoinRequestWithContext(ctx context.Context, userID string, accept bool) error {
	return c.client.ResolveChatPendingWithContext(ctx, c.ID, userID, accept, false)
}

// ResolveAllJoinRequest resolve all chat join requests
//
// ResolveAllJoinRequests uses context.Background internally; to specify the context, use
// ResolveAllJoinRequestsWithContext.
func (c *Chat) ResolveAllJoinRequests(accept bool) error {
	return c.ResolveAllJoinRequestsWithContext(context.Background(), accept)
}

// ResolveAllJoinRequest resolve all chat join requests
func (c *Chat) ResolveAllJoinRequestsWithContext(ctx context.Context, accept bool) error {
	return c.client.ResolveChatPendingWithContext(ctx, c.ID, "", accept, true)
}

// SetTitle changes chat title
//
// SetTitle uses context.Background internally; to specify the context, use
// SetTitleWithContext.
func (c *Chat) SetTitle(title string) error {
	return c.SetTitleWithContext(context.Background(), title)
}

// SetTitle changes chat title
func (c *Chat) SetTitleWithContext(ctx context.Context, title string) error {
	return c.client.SetChatTitleWithContext(ctx, c.ID, title)
}

// SetAbout changes chat about
//
// SetAbout uses context.Background internally; to specify the context, use
// SetAboutWithContext.
func (c *Chat) SetAbout(about string) error {
	return c.SetAboutWithContext(context.Background(), about)
}

// SetAbout changes chat about
func (c *Chat) SetAboutWithContext(ctx context.Context, about string) error {
	return c.client.SetChatAboutWithContext(ctx, c.ID, about)
}

// SetRules changes chat rules
//
// SetRules uses context.Background internally; to specify the context, use
// SetRulesWithContext.
func (c *Chat) SetRules(rules string) error {
	return c.SetRulesWithContext(context.Background(), rules)
}

// SetRules changes chat rules
func (c *Chat) SetRulesWithContext(ctx context.Context, rules string) error {
	return c.client.SetChatRulesWithContext(ctx, c.ID, rules)
}
