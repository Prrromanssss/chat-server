package tests

import (
	"context"
	"testing"

	"github.com/Prrromanssss/platform_common/pkg/db"
	dbMocks "github.com/Prrromanssss/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/Prrromanssss/chat-server/internal/model"
	"github.com/Prrromanssss/chat-server/internal/repository"
	repositoryMocks "github.com/Prrromanssss/chat-server/internal/repository/mocks"
	chatService "github.com/Prrromanssss/chat-server/internal/service/chat"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()

	type (
		chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
		logRepositoryMockFunc  func(mc *minimock.Controller) repository.LogRepository
		txManagerMockFunc      func(f func(context.Context) error, mc *minimock.Controller) db.TxManager
	)

	type args struct {
		ctx context.Context
		req model.DeleteChatParams
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		ErrUserRepository = errors.New("user repository error")
		ErrLogRepository  = errors.New("log repository error")

		req = model.DeleteChatParams{
			ChatID: id,
		}

		logApiReq = model.CreateAPILogParams{
			Method:      "Delete",
			RequestData: req,
		}
	)

	tests := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(nil)
				mock.DeleteChatMock.Expect(ctx, req).Return(nil)

				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMocks.NewLogRepositoryMock(mc)
				mock.CreateAPILogMock.Expect(ctx, logApiReq).Return(nil)

				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return mock
			},
		},
		{
			name: "user repository error in DeleteChat",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrUserRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(nil)
				mock.DeleteChatMock.Expect(ctx, req).Return(ErrUserRepository)

				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMocks.NewLogRepositoryMock(mc)

				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return mock
			},
		},
		{
			name: "user repository error in UnlinkParticipantsFromChatParams",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrUserRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(ErrUserRepository)

				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMocks.NewLogRepositoryMock(mc)

				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return mock
			},
		},
		{
			name: "log repository error",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrLogRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(nil)
				mock.DeleteChatMock.Expect(ctx, req).Return(nil)

				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMocks.NewLogRepositoryMock(mc)
				mock.CreateAPILogMock.Expect(ctx, logApiReq).Return(ErrLogRepository)

				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			logRepositoryMock := tt.logRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				txErr := chatRepositoryMock.DeleteChat(ctx, req)
				if txErr != nil {
					return txErr
				}

				txErr = logRepositoryMock.CreateAPILog(ctx, model.CreateAPILogParams{
					Method:      "Delete",
					RequestData: req,
				})
				if txErr != nil {
					return txErr
				}

				return nil
			}, mc)

			service := chatService.NewService(chatRepositoryMock, logRepositoryMock, txManagerMock)

			err := service.DeleteChat(tt.args.ctx, tt.args.req)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
