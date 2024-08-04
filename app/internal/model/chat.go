package model

import "time"

// CreateChatResponse represents the response after creating a chat, including the ChatID.
type CreateChatResponse struct {
	ChatID int64
}

type CreateUsersForChatParams struct {
	Emails []string
}

type CreateUsersForChatResponse struct {
	UserIDs []int64
}

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
	From   string    // Sender of the message
	Text   string    // Message content
	SentAt time.Time // Timestamp of the message
}

type UnlinkParticipantsFromChatParams struct {
	ChatID int64
}
