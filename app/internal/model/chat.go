package models

import "time"

// SendMessageParams holds the data for sending a message.
type SendMessageParams struct {
	From   string    // Sender of the message
	Text   string    // Message content
	SentAt time.Time // Timestamp of the message
}

// LinkParticipantsToChat links a user to a chat.
type LinkParticipantsToChat struct {
	ChatID int64 `db:"chat_id"` // Chat identifier
	UserID int64 `db:"user_id"` // User identifier
}
