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

// NewPGRepo creates a new instance of chatPGRepo with the provided SQL database connection.
func NewPGRepo(db *sqlx.DB) repository.ChatRepository {
	return &chatPGRepo{db: db}
}

// CreateChat creates a new chat and links participants to it.
func (p *chatPGRepo) CreateChat(ctx context.Context, emails []string) (chatID int64, err error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, errors.Wrap(err, "chatPGRepo.CreateChat.BeginTxx")
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			log.Println(rollbackErr)
		}
	}()

	err = tx.GetContext(ctx, &chatID, queryCreateChat)
	if err != nil {
		return 0, errors.Wrap(err, "chatPGRepo.CreateChat.GetContext.queryCreateChat")
	}

	userIDs := make([]int64, len(emails))

	stmt, err := tx.PrepareContext(ctx, queryCreateUser)
	if err != nil {
		return 0, errors.Wrap(err, "chatPGRepo.CreateChat.PrepareContext.queryCreateUser")
	}

	defer func() {
		if stmtErr := stmt.Close(); stmtErr != nil {
			log.Println(stmtErr)
		}
	}()

	for _, email := range emails {
		var userID int64
		err = stmt.QueryRowContext(ctx, email).Scan(&userID)
		if err != nil {
			return 0, errors.Wrap(err, "chatPGRepo.CreateUsers.QueryRowContext.Scan")
		}
		userIDs = append(userIDs, userID)
	}

	params := make([]models.LinkParticipantsToChat, len(userIDs))
	for i, userID := range userIDs {
		params[i] = models.LinkParticipantsToChat{UserID: userID, ChatID: chatID}
	}

	_, err = tx.NamedExecContext(ctx, queryLinkParticipantsToChat, params)
	if err != nil {
		return 0, errors.Wrap(err, "chatPGRepo.CreateChat.queryLinkParticipantsToChat")
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Wrap(err, "chatPGRepo.CreateChat.Commit")
	}

	return chatID, nil
}

// DeleteChat removes a chat and unlinks its participants.
func (p *chatPGRepo) DeleteChat(ctx context.Context, chatID int64) (err error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "chatPGRepo.DeleteChat.BeginTxx")
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			log.Println(rollbackErr)
		}
	}()

	_, err = tx.ExecContext(ctx, queryUnlinkParticipantsFromChat, chatID)
	if err != nil {
		return errors.Wrapf(err, "chatPGRepo.DeleteChat.ExecContext.queryUnlinkParticipantsFromChat(chatID: %d)", chatID)
	}

	_, err = tx.ExecContext(ctx, queryDeleteChat, chatID)
	if err != nil {
		return errors.Wrapf(err, "chatPGRepo.DeleteChat.ExecContext.queryDeleteChat(chatID: %d)", chatID)
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "chatPGRepo.DeleteChat.Commit")
	}

	return nil
}

// SendMessage sends a message to a chat.
func (p *chatPGRepo) SendMessage(ctx context.Context, params models.SendMessageParams) (err error) {
	_, err = p.db.ExecContext(ctx, querySendMessage, params.From, params.Text, params.SentAt)
	if err != nil {
		return errors.Wrapf(err, "chatPGRepo.SendMessage.ExecContext.querySendMessage(From: %s)", params.From)
	}

	return nil
}
