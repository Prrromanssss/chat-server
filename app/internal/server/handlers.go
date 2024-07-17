package server

import (
	"context"

	deliveryGRPC "github.com/Prrromanssss/chat-server/internal/delivery/grpc"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

func (s *Server) MapHandlers(ctx context.Context) error {
	GRPCHandlers := deliveryGRPC.NewGRPCHandlers()
	pb.RegisterChatV1Server(s.grpc, GRPCHandlers)

	return nil
}
