package repository

import (
	"context"

	"github.com/Prrromanssss/chat-server/internal/models"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, userIDs []int64) (chatID int64, err error)
	DeleteChat(ctx context.Context, chatID int64) (err error)
	SendMessage(ctx context.Context, params models.SendMessageParams) (err error)
}
