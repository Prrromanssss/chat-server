package chat

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/Prrromanssss/chat-server/internal/client/db"
	"github.com/Prrromanssss/chat-server/internal/model"
	"github.com/Prrromanssss/chat-server/internal/repository"
	"github.com/Prrromanssss/chat-server/internal/repository/chat/converter"
	modelRepo "github.com/Prrromanssss/chat-server/internal/repository/chat/model"
)

type chatPGRepo struct {
	db db.Client
}

// NewRepository creates a new instance of chatPGRepo with the provided SQL database connection.
func NewRepository(db db.Client) repository.ChatRepository {
	return &chatPGRepo{
		db: db,
	}
}

// CreateChat creates a new chat and links participants to it.
func (p *chatPGRepo) CreateChat(ctx context.Context) (resp model.CreateChatResponse, err error) {
	log.Infof("chatPGRepo.CreateChat, params")

	var respRepo modelRepo.CreateChatResponse

	q := db.Query{
		Name:     "chatPGRepo.CreateChat",
		QueryRaw: queryCreateChat,
	}

	err = p.db.DB().ScanOneContext(ctx, &respRepo, q)
	if err != nil {
		err = errors.Wrap(err, "Cannot create chat")
		return
	}

	return converter.ConvertCreateChatResponseFromRepoToService(respRepo), nil
}

// CreateUsersForChat creates users for the chat based on the provided email list and returns their IDs.
func (p *chatPGRepo) CreateUsersForChat(
	ctx context.Context,
	params model.CreateUsersForChatParams,
) (resp model.CreateUsersForChatResponse, err error) {
	log.Infof("chatPGRepo.CreateUsersForChat, params: %+v", params)

	paramsRepo := converter.ConvertCreateUsersForChatParamsFromServiceToRepo(params)

	userIDs := make([]int64, len(paramsRepo.Emails))

	for i, email := range paramsRepo.Emails {
		var userID int64

		q := db.Query{
			Name:     "chatPGRepo.CreateUsersForChat",
			QueryRaw: queryCreateUser,
		}

		err = p.db.DB().ScanOneContext(ctx, &userID, q, email)
		if err != nil {
			err = errors.Wrapf(err, "Cannot create user for chat(userID: %v)", userID)
			return
		}

		userIDs[i] = userID
	}

	return converter.ConvertCreateUsersForChatResponseFromRepoToService(
		modelRepo.CreateUsersForChatResponse{
			UserIDs: userIDs,
		},
	), nil
}

// LinkParticipantsToChat links participants to a chat by adding their IDs to the chat participants list.
func (p *chatPGRepo) LinkParticipantsToChat(
	ctx context.Context,
	params model.LinkParticipantsToChatParams,
) (err error) {
	log.Infof("chatPGRepo.LinkParticipantsToChat, params: %+v", params)

	paramsRepo := converter.ConvertLinkParticipantsToChatParamsFromServiceToRepo(params)

	batch := &pgx.Batch{}

	for _, userID := range paramsRepo.UserIDs {
		batch.Queue(queryLinkParticipantsToChat, paramsRepo.ChatID, userID)
	}

	br := p.db.DB().SendBatchContext(ctx, batch)

	err = br.Close()
	if err != nil {
		err = errors.Wrapf(
			err,
			"Cannot close batch for chat(chatID: %d)",
			paramsRepo.ChatID,
		)
		return
	}

	return nil
}

// UnlinkParticipantsFromChat removes participants from a chat based on the provided parameters.
func (p *chatPGRepo) UnlinkParticipantsFromChat(
	ctx context.Context,
	params model.UnlinkParticipantsFromChatParams,
) (err error) {
	log.Infof("chatPGRepo.UnlinkParticipantsFromChat, params: %+v", params)

	paramsRepo := converter.ConvertUnlinkParticipantsFromChatParamsFromServiceToRepo(params)

	q := db.Query{
		Name:     "chatPGRepo.UnlinkParticipantsFromChat",
		QueryRaw: queryUnlinkParticipantsFromChat,
	}

	_, err = p.db.DB().ExecContext(ctx, q, paramsRepo.ChatID)
	if err != nil {
		err = errors.Wrapf(
			err,
			"Cannot unlink participants from chat(chatID: %d)",
			paramsRepo.ChatID,
		)
		return
	}

	return nil
}

// DeleteChat removes a chat and unlinks its participants based on the provided chat ID.
func (p *chatPGRepo) DeleteChat(ctx context.Context, params model.DeleteChatParams) (err error) {
	log.Infof("chatPGRepo.DeleteChat, params: %+v", params)

	paramsRepo := converter.ConvertDeleteChatParamsFromServiceToRepo(params)

	q := db.Query{
		Name:     "chatPGRepo.DeleteChat",
		QueryRaw: queryDeleteChat,
	}

	_, err = p.db.DB().ExecContext(ctx, q, paramsRepo.ChatID)
	if err != nil {
		err = errors.Wrapf(err, "Cannot delete chat(chatID: %d)", paramsRepo.ChatID)
		return
	}

	return nil
}

// SendMessage sends a message to a chat.
func (p *chatPGRepo) SendMessage(ctx context.Context, params model.SendMessageParams) (err error) {
	log.Infof("chatPGRepo.SendMessage, params: %+v", params)

	paramsRepo := converter.ConvertSendMessageParamsFromServiceToRepo(params)

	q := db.Query{
		Name:     "chatPGRepo.SendMessage",
		QueryRaw: querySendMessage,
	}

	_, err = p.db.DB().ExecContext(ctx, q, paramsRepo.From, paramsRepo.Text, paramsRepo.SentAt)
	if err != nil {
		err = errors.Wrapf(err, "Cannot send message (from: %s)", paramsRepo.From)
		return
	}

	return nil
}

// CreateAPILog creates log in database of every api action.
func (p *chatPGRepo) CreateAPILog(
	ctx context.Context,
	params model.CreateAPILogParams,
) (err error) {
	log.Infof("chatPGRepo.CreateAPILog, params: %+v", params)

	paramsRepo := converter.ConvertCreateAPILogParamsFromServiceToRepo(params)

	q := db.Query{
		Name:     "chatPGRepo.CreateAPILog",
		QueryRaw: queryCreateAPILog,
	}

	_, err = p.db.DB().ExecContext(ctx, q, paramsRepo.Method, paramsRepo.RequestData, paramsRepo.ResponseData)
	if err != nil {
		return errors.Wrapf(
			err,
			"Cannot create api log for chat",
		)
	}

	return nil
}
