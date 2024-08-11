package chat

import (
	"context"
	"fmt"

	"github.com/Prrromanssss/platform_common/pkg/db"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"

	"github.com/Prrromanssss/chat-server/internal/model"
	"github.com/Prrromanssss/chat-server/internal/repository"
	"github.com/Prrromanssss/chat-server/internal/service"
)

type chatService struct {
	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService creates a new instance of chatService with the provided ChatRepository and TxManager.
func NewService(
	chatRepository repository.ChatRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) service.ChatService {
	return &chatService{
		chatRepository: chatRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}

// CreateChat handles the creation of a new chat and links participants to it within a transaction.
// It also logs the request and response data for auditing purposes.
func (s *chatService) CreateChat(
	ctx context.Context,
	params model.CreateChatParams,
) (resp model.CreateChatResponse, err error) {
	log.Infof("chatService.CreateChat")

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		resp, txErr = s.chatRepository.CreateChat(ctx)
		if txErr != nil {
			return txErr
		}

		fmt.Printf("ChatID: %d\n", resp.ChatID)

		// Create users for the chat and get their IDs
		usersResp, txErr := s.chatRepository.CreateUsersForChat(ctx, model.CreateUsersForChatParams(params))
		if txErr != nil {
			return txErr
		}

		fmt.Printf("Users: %v", usersResp.UserIDs)

		// Link the created users to the new chat
		txErr = s.chatRepository.LinkParticipantsToChat(ctx, model.LinkParticipantsToChatParams{
			ChatID:  resp.ChatID,
			UserIDs: usersResp.UserIDs,
		})
		if txErr != nil {
			return txErr
		}

		txErr = s.logRepository.CreateAPILog(ctx, model.CreateAPILogParams{
			Method:       "Create",
			RequestData:  params,
			ResponseData: resp,
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return model.CreateChatResponse{}, errors.Wrapf(err, "Transaction failed")
	}

	return resp, nil
}

// DeleteChat handles the deletion of a chat and unlinks participants from it within a transaction.
// It also logs the request data for auditing purposes.
func (s *chatService) DeleteChat(ctx context.Context, params model.DeleteChatParams) (err error) {
	log.Infof("chatService.DeleteChat, params: %v", params)

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		txErr := s.chatRepository.UnlinkParticipantsFromChat(ctx, model.UnlinkParticipantsFromChatParams(params))
		if txErr != nil {
			return txErr
		}

		txErr = s.chatRepository.DeleteChat(ctx, params)
		if txErr != nil {
			return txErr
		}

		txErr = s.logRepository.CreateAPILog(ctx, model.CreateAPILogParams{
			Method:      "Delete",
			RequestData: params,
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		err = errors.Wrapf(err, "Transaction failed")
		return
	}

	return nil
}

// SendMessage handles sending a message within a transaction.
// It also logs the request data for auditing purposes.
func (s *chatService) SendMessage(ctx context.Context, params model.SendMessageParams) (err error) {
	log.Infof("chatService.SendMessage, params: %v", params)

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		txErr := s.chatRepository.SendMessage(ctx, params)
		if txErr != nil {
			return txErr
		}

		txErr = s.logRepository.CreateAPILog(ctx, model.CreateAPILogParams{
			Method:      "SendMessage",
			RequestData: params,
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		err = errors.Wrapf(err, "Transaction failed")
		return
	}

	return nil
}
