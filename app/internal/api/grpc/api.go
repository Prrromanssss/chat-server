package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Prrromanssss/chat-server/internal/models"
	"github.com/Prrromanssss/chat-server/internal/repository"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

// GRPCHandlers implements the gRPC server for chat operations.
// It uses a ChatRepository to interact with chat data.
type GRPCHandlers struct {
	pb.UnimplementedChatV1Server                           // Embeds the unimplemented server for backward compatibility.
	repo                         repository.ChatRepository // Repository instance for chat data operations.
}

// NewGRPCHandlers creates a new instance of GRPCHandlers with the provided ChatRepository.
func NewGRPCHandlers(repo repository.ChatRepository) pb.ChatV1Server {
	return &GRPCHandlers{repo: repo}
}

// Create handles the RPC call to create a new chat.
// It takes a CreateRequest, creates a chat, and returns a CreateResponse with the new chat ID.
func (h *GRPCHandlers) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("rpc Create, request: %+v", req)

	chatID, err := h.repo.CreateChat(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Id: chatID,
	}, nil
}

// Delete handles the RPC call to delete an existing chat.
// It takes a DeleteRequest, deletes the chat, and returns an empty response.
func (h *GRPCHandlers) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("rpc Delete, request: %+v", req)

	err := h.repo.DeleteChat(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SendMessage handles the RPC call to send a message to a chat.
// It takes a SendMessageRequest, sends the message, and returns an empty response.
func (h *GRPCHandlers) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("rpc SendMessage, request: %+v", req)

	err := h.repo.SendMessage(ctx, models.SendMessageParams{
		From:   req.From,
		Text:   req.Text,
		SentAt: req.Timestamp.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
