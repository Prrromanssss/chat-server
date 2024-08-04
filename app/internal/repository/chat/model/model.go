package model

import (
	"database/sql"
	"time"
)

// CreateUsersForChatParams holds the list of email addresses for creating users in a chat.
type CreateUsersForChatParams struct {
	Emails []string `db:"emails"`
}

// CreateChatResponse represents the response after creating a chat, including the ChatID.
type CreateChatResponse struct {
	ChatID int64 `db:"id"`
}

// CreateUsersForChatResponse represents the response after creating users for a chat.
type CreateUsersForChatResponse struct {
	UserIDs []int64
}

// DeleteChatParams holds the ID of the chat to be deleted.
type DeleteChatParams struct {
	ChatID int64 `db:"id"`
}

// SendMessageParams holds the data for sending a message.
type SendMessageParams struct {
	From   string    `db:"sender"`       // Sender of the message
	Text   string    `db:"message_text"` // Message content
	SentAt time.Time `db:"sent_at"`      // Timestamp of the message
}

// LinkParticipantsToChatParams holds the data for linking users to a chat.
type LinkParticipantsToChatParams struct {
	ChatID  int64   `db:"chat_id"`
	UserIDs []int64 `db:"user_ids"`
}

// UnlinkParticipantsFromChatParams holds the data for unlinking users from a chat.
type UnlinkParticipantsFromChatParams struct {
	ChatID int64 `db:"chat_id"`
}

// CreateAPILogParams holds the parameters for logging API actions related to user creation.
type CreateAPILogParams struct {
	Method       string         `db:"action_type"`
	RequestData  string         `db:"request_data"`
	ResponseData sql.NullString `db:"response_data"`
}
