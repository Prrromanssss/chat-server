package repository

import (
	"context"

	"github.com/Prrromanssss/chat-server/internal/model"
)

// ChatRepository defines methods for managing chat operations.
type ChatRepository interface {
	// CreateChat creates a new chat with the given user emails and returns the chat ID.
	CreateChat(ctx context.Context) (resp model.CreateChatResponse, err error)

	// CreateUsersForChat creates users for a chat based on the provided parameters
	// and returns the response with the created user IDs.
	CreateUsersForChat(
		ctx context.Context,
		params model.CreateUsersForChatParams,
	) (resp model.CreateUsersForChatResponse, err error)

	// LinkParticipantsToChat links users to a chat based on the provided parameters.
	LinkParticipantsToChat(ctx context.Context, params model.LinkParticipantsToChatParams) (err error)

	// UnlinkParticipantsFromChat unlinks users from a chat based on the provided parameters.
	UnlinkParticipantsFromChat(ctx context.Context, params model.UnlinkParticipantsFromChatParams) (err error)

	// DeleteChat removes a chat identified by its chat ID.
	DeleteChat(ctx context.Context, params model.DeleteChatParams) (err error)

	// SendMessage sends a message with the specified parameters.
	SendMessage(ctx context.Context, params model.SendMessageParams) (err error)

	// CreateAPILog creates log in database of every api action and returns any error..
	CreateAPILog(ctx context.Context, params model.CreateAPILogParams) (err error)
}
