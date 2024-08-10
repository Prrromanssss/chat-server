package tests

import (
	"context"
	"encoding/json"
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

		ErrRepository    = errors.New("repository error")
		ErrRepositoryLog = errors.New("repository error in CreateAPILog")

		req = model.DeleteChatParams{
			ChatID: id,
		}
	)

	requestData, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	logApiReq := model.CreateAPILogParams{
		Method:      "Delete",
		RequestData: string(requestData),
	}

	tests := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMockFunc
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
			name: "repository error case in DeleteChat",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(nil)
				mock.DeleteChatMock.Expect(ctx, req).Return(ErrRepository)

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
			name: "repository error case in UnlinkParticipantsFromChatParams",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(ErrRepository)

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
			name: "repository error in CreateAPILog",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrRepositoryLog,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.UnlinkParticipantsFromChatMock.Expect(ctx, model.UnlinkParticipantsFromChatParams(req)).
					Return(nil)
				mock.DeleteChatMock.Expect(ctx, req).Return(nil)
				mock.CreateAPILogMock.Expect(ctx, logApiReq).Return(ErrRepositoryLog)

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
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				txErr := chatRepositoryMock.DeleteChat(ctx, req)
				if txErr != nil {
					return txErr
				}

				requestData, txErr := json.Marshal(req)
				if txErr != nil {
					return txErr
				}

				txErr = chatRepositoryMock.CreateAPILog(ctx, model.CreateAPILogParams{
					Method:      "Delete",
					RequestData: string(requestData),
				})

				if txErr != nil {
					return txErr
				}

				return nil
			}, mc)
			service := chatService.NewService(chatRepositoryMock, txManagerMock)

			err := service.DeleteChat(tt.args.ctx, tt.args.req)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
