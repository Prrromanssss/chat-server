package server

import (
	"context"

	deliveryGRPC "github.com/Prrromanssss/chat-server/internal/api/grpc"
	"github.com/Prrromanssss/chat-server/internal/repository/chat"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

// MapHandlers initializes and registers the gRPC handlers with the server.
// It sets up the repository and maps the gRPC handlers to the server.
func (s *Server) MapHandlers(ctx context.Context) error {
	// Initialize repository
	userRepo := chat.NewPGRepo(s.pgDB)

	// Create and register gRPC handlers
	GRPCHandlers := deliveryGRPC.NewGRPCHandlers(userRepo)
	pb.RegisterChatV1Server(s.grpc, GRPCHandlers)

	return nil
}
