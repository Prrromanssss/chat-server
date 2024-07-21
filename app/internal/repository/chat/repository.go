package chat

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Prrromanssss/chat-server/internal/models"
	"github.com/Prrromanssss/chat-server/internal/repository"
)

type chatPGRepo struct {
	db *sqlx.DB
}

func NewPGRepo(db *sqlx.DB) repository.ChatRepository {
	return &chatPGRepo{db: db}
}

func (p *chatPGRepo) CreateChat(ctx context.Context, userIDs []int64) (chatID int64, err error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, errors.Wrap(
			err,
			"chatPGRepo.CreateChat.BeginTxx",
		)
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			log.Println(rollbackErr)
		}
	}()

	err = tx.GetContext(ctx, &chatID, queryCreateChat)
	if err != nil {
		return 0, errors.Wrap(
			err,
			"chatPGRepo.CreateChat.GetContext.queryCreateUser",
		)
	}

	params := make([]models.LinkParticipantsToChat, len(userIDs))
	for i, userID := range userIDs {
		params[i] = models.LinkParticipantsToChat{UserID: userID, ChatID: chatID}
	}

	_, err = tx.NamedExecContext(ctx, queryLinkParticipantsToChat, params)
	if err != nil {
		return 0, errors.Wrap(
			err,
			`chatPGRepo.CreateChat.queryLinkParticipantsToChat 
			Unable to link participants to chat`,
		)
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Wrap(
			err,
			"chatPGRepo.CreateChat.Commit",
		)
	}

	return chatID, nil
}

func (p *chatPGRepo) DeleteChat(ctx context.Context, chatID int64) (err error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(
			err,
			"chatPGRepo.DeleteChat.BeginTxx",
		)
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			log.Println(rollbackErr)
		}
	}()

	_, err = tx.ExecContext(ctx, queryUnlinkParticipantsFromChat, chatID)
	if err != nil {
		return errors.Wrapf(
			err,
			"chatPGRepo.DeleteChat.ExecContext.queryUnlinkParticipantsFromChat(chatID: %d)",
			chatID,
		)
	}

	_, err = tx.ExecContext(ctx, queryDeleteChat, chatID)
	if err != nil {
		return errors.Wrapf(
			err,
			"chatPGRepo.DeleteChat.ExecContext.queryDeleteChat(chatID: %d)",
			chatID,
		)
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(
			err,
			"chatPGRepo.DeleteChat.Commit",
		)
	}

	return nil
}

func (p *chatPGRepo) SendMessage(ctx context.Context, params models.SendMessageParams) (err error) {
	_, err = p.db.ExecContext(ctx, querySendMessage, params.From, params.Text, params.SentAt)
	if err != nil {
		return errors.Wrapf(
			err,
			"chatPGRepo.SendMessage.ExecContext.querySendMessage(From: %s)",
			params.From,
		)
	}

	return nil
}