package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	chatAPI "github.com/Prrromanssss/chat-server/internal/api/grpc/chat"
	"github.com/Prrromanssss/chat-server/internal/model"
	"github.com/Prrromanssss/chat-server/internal/service"
	serviceMocks "github.com/Prrromanssss/chat-server/internal/service/mocks"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *pb.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = gofakeit.Int64()
		email1 = gofakeit.Email()
		email2 = gofakeit.Email()

		ErrService = errors.New("service error")

		req = &pb.CreateRequest{
			Emails: []string{email1, email2},
		}

		serviceParams = model.CreateChatParams{
			Emails: []string{email1, email2},
		}

		resp = &pb.CreateResponse{
			Id: id,
		}

		serviceResp = model.CreateChatResponse{
			ChatID: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *pb.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, serviceParams).Return(serviceResp, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  ErrService,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, serviceParams).Return(model.CreateChatResponse{}, ErrService)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chatAPI.NewGRPCHandlers(chatServiceMock)

			resp, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}
}
