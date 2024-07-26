package repository

import (
	"context"

	"github.com/Prrromanssss/chat-server/internal/models"
)

// ChatRepository defines methods for managing chat operations.
type ChatRepository interface {
	// CreateChat creates a new chat with the given user IDs and returns the chat ID.
	CreateChat(ctx context.Context, userIDs []int64) (chatID int64, err error)

	// DeleteChat removes a chat identified by its chat ID.
	DeleteChat(ctx context.Context, chatID int64) (err error)

	// SendMessage sends a message with the specified parameters.
	SendMessage(ctx context.Context, params models.SendMessageParams) (err error)
}
