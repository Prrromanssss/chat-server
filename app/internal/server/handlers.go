package server

import (
	"context"

	deliveryGRPC "github.com/Prrromanssss/chat-server/internal/delivery/grpc"
	"github.com/Prrromanssss/chat-server/internal/repository/chat"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

func (s *Server) MapHandlers(ctx context.Context) error {
	// repos
	userRepo := chat.NewPGRepo(s.pgDB)

	// handlers
	GRPCHandlers := deliveryGRPC.NewGRPCHandlers(userRepo)

	pb.RegisterChatV1Server(s.grpc, GRPCHandlers)

	return nil
}
