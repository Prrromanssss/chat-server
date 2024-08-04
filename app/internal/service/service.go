package service

import (
	"context"

	"github.com/Prrromanssss/chat-server/internal/model"
)

type ChatService interface {
	// CreateChat creates a new chat with the given user emails and returns the chat ID.
	CreateChat(ctx context.Context, params model.CreateChatParams) (resp model.CreateChatResponse, err error)

	// DeleteChat removes a chat identified by its chat ID.
	DeleteChat(ctx context.Context, params model.DeleteChatParams) (err error)

	// SendMessage sends a message with the specified parameters.
	SendMessage(ctx context.Context, params model.SendMessageParams) (err error)
}
