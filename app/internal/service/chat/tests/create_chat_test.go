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

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type (
		chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
		txManagerMockFunc      func(f func(context.Context) error, mc *minimock.Controller) db.TxManager
	)

	type args struct {
		ctx context.Context
		req model.CreateChatParams
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID  = gofakeit.Int64()
		id1     = gofakeit.Int64()
		id2     = gofakeit.Int64()
		email1  = gofakeit.Email()
		email2  = gofakeit.Email()
		userIDs = []int64{id1, id2}

		ErrRepository    = errors.New("repository error")
		ErrRepositoryLog = errors.New("repository error in CreateAPILog")

		req = model.CreateChatParams{
			Emails: []string{email1, email2},
		}

		resp = model.CreateChatResponse{
			ChatID: chatID,
		}

		usersResp = model.CreateUsersForChatResponse{
			UserIDs: userIDs,
		}
	)

	requestData, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	responseData, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
	}

	responseDataString := string(responseData)

	logApiReq := model.CreateAPILogParams{
		Method:       "Create",
		RequestData:  string(requestData),
		ResponseData: &responseDataString,
	}

	tests := []struct {
		name               string
		args               args
		want               model.CreateChatResponse
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
			want: resp,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx).Return(resp, nil)
				mock.CreateUsersForChatMock.Expect(ctx, model.CreateUsersForChatParams(req)).
					Return(usersResp, nil)
				mock.LinkParticipantsToChatMock.Expect(ctx, model.LinkParticipantsToChatParams{
					ChatID:  chatID,
					UserIDs: userIDs,
				}).Return(nil)
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
			name: "repository error case in CreateChat",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: model.CreateChatResponse{},
			err:  ErrRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx).Return(model.CreateChatResponse{}, ErrRepository)

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
			name: "repository error case in CreateUsersForChat",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: model.CreateChatResponse{},
			err:  ErrRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx).Return(resp, nil)
				mock.CreateUsersForChatMock.Expect(ctx, model.CreateUsersForChatParams(req)).
					Return(model.CreateUsersForChatResponse{}, ErrRepository)

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
			name: "repository error case in LinkParticipantsToChat",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: model.CreateChatResponse{},
			err:  ErrRepository,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx).Return(resp, nil)
				mock.CreateUsersForChatMock.Expect(ctx, model.CreateUsersForChatParams(req)).
					Return(usersResp, nil)
				mock.LinkParticipantsToChatMock.Expect(ctx, model.LinkParticipantsToChatParams{
					ChatID:  chatID,
					UserIDs: userIDs,
				}).Return(ErrRepository)

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
			want: model.CreateChatResponse{},
			err:  ErrRepositoryLog,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx).Return(resp, nil)
				mock.CreateUsersForChatMock.Expect(ctx, model.CreateUsersForChatParams(req)).
					Return(usersResp, nil)
				mock.LinkParticipantsToChatMock.Expect(ctx, model.LinkParticipantsToChatParams{
					ChatID:  chatID,
					UserIDs: userIDs,
				}).Return(nil)
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
				resp, txErr := chatRepositoryMock.CreateChat(ctx)
				if txErr != nil {
					return txErr
				}

				usersResp, txErr := chatRepositoryMock.CreateUsersForChat(ctx, model.CreateUsersForChatParams(req))
				if txErr != nil {
					return txErr
				}

				txErr = chatRepositoryMock.LinkParticipantsToChat(ctx, model.LinkParticipantsToChatParams{
					ChatID:  resp.ChatID,
					UserIDs: usersResp.UserIDs,
				})
				if txErr != nil {
					return txErr
				}

				requestData, txErr := json.Marshal(req)
				if txErr != nil {
					return txErr
				}

				responseData, err := json.Marshal(resp)
				if err != nil {
					t.Error(err)
				}

				responseDataString := string(responseData)

				txErr = chatRepositoryMock.CreateAPILog(ctx, model.CreateAPILogParams{
					Method:       "Create",
					RequestData:  string(requestData),
					ResponseData: &responseDataString,
				})

				if txErr != nil {
					return txErr
				}

				return nil
			}, mc)
			service := chatService.NewService(chatRepositoryMock, txManagerMock)

			resp, err := service.CreateChat(tt.args.ctx, tt.args.req)
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.want, resp)
		})
	}
}
