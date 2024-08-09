package app

import (
	"context"
	"log"

	"github.com/Prrromanssss/chat-server/config"
	chatAPI "github.com/Prrromanssss/chat-server/internal/api/grpc/chat"
	"github.com/Prrromanssss/chat-server/internal/client/db"
	"github.com/Prrromanssss/chat-server/internal/client/db/pg"
	"github.com/Prrromanssss/chat-server/internal/client/db/transaction"
	"github.com/Prrromanssss/chat-server/internal/repository"
	chatRepository "github.com/Prrromanssss/chat-server/internal/repository/chat"
	"github.com/Prrromanssss/chat-server/internal/service"
	chatService "github.com/Prrromanssss/chat-server/internal/service/chat"
	"github.com/Prrromanssss/chat-server/pkg/closer"
)

type serviceProvider struct {
	cfg *config.Config

	db        db.Client
	txManager db.TxManager

	chatRepository repository.ChatRepository
	chatService    service.ChatService
	chatAPI        *chatAPI.GRPCHandlers
}

func newServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{
		cfg: cfg,
	}
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.db == nil {
		cl, err := pg.New(ctx, s.cfg.Postgres.DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.db = cl
	}

	return s.db
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatAPI(ctx context.Context) *chatAPI.GRPCHandlers {
	if s.chatAPI == nil {
		s.chatAPI = chatAPI.NewGRPCHandlers(s.ChatService(ctx))
	}

	return s.chatAPI
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
