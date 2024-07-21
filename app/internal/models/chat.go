package models

import "time"

type SendMessageParams struct {
	From   string
	Text   string
	SentAt time.Time
}

type LinkParticipantsToChat struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}
