package model

import "time"

// CreateChatParams contains the parameters for creating users in a chat.
type CreateChatParams struct {
	Emails []string
}

// CreateChatResponse represents the response after creating a chat, including the ChatID.
type CreateChatResponse struct {
	ChatID int64
}

// CreateUsersForChatParams contains the parameters for creating users in a chat.
type CreateUsersForChatParams struct {
	Emails []string
}

// CreateUsersForChatResponse represents the response after creating users for a chat.
type CreateUsersForChatResponse struct {
	UserIDs []int64
}

// LinkParticipantsToChatParams contains the parameters for linking users to a chat.
type LinkParticipantsToChatParams struct {
	ChatID  int64
	UserIDs []int64
}

// DeleteChatParams holds the ID of the chat to be deleted.
type DeleteChatParams struct {
	ChatID int64
}

// SendMessageParams holds the data for sending a message.
type SendMessageParams struct {
	From   string
	Text   string
	SentAt time.Time
}

// UnlinkParticipantsFromChatParams holds the ID of the chat from which users will be unlinked.
type UnlinkParticipantsFromChatParams struct {
	ChatID int64
}

// CreateAPILogParams holds the parameters for logging API actions related to user creation.
type CreateAPILogParams struct {
	Method       string
	RequestData  interface{}
	ResponseData interface{}
}
