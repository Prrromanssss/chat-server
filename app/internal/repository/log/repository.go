package log

import (
	"context"

	"github.com/Prrromanssss/platform_common/pkg/db"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"

	"github.com/Prrromanssss/chat-server/internal/model"
	"github.com/Prrromanssss/chat-server/internal/repository"
	"github.com/Prrromanssss/chat-server/internal/repository/log/converter"
)

type logPGRepo struct {
	db db.Client
}

// NewRepository creates a new instance of logPGRepo with the provided database connection.
func NewRepository(db db.Client) repository.LogRepository {
	return &logPGRepo{db: db}
}

// CreateAPILog creates log in database of every api action.
func (p *logPGRepo) CreateAPILog(
	ctx context.Context,
	params model.CreateAPILogParams,
) (err error) {
	log.Infof("logPGRepo.CreateAPILog, params: %+v", params)

	paramsRepo, err := converter.ConvertCreateAPILogParamsFromServiceToRepo(params)
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "logPGRepo.CreateAPILog",
		QueryRaw: queryCreateAPILog,
	}

	_, err = p.db.DB().ExecContext(ctx, q, paramsRepo.Method, paramsRepo.RequestData, paramsRepo.ResponseData)
	if err != nil {
		return errors.Wrapf(
			err,
			"Cannot create api log for user",
		)
	}

	return nil
}
