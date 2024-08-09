package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Prrromanssss/chat-server/internal/converter"
	"github.com/Prrromanssss/chat-server/internal/service"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

// GRPCHandlers implements the gRPC server for chat operations.
// It uses a ChatRepository to interact with chat data.
type GRPCHandlers struct {
	pb.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewGRPCHandlers creates a new instance of GRPCHandlers with the provided ChatRepository.
func NewGRPCHandlers(chatService service.ChatService) *GRPCHandlers {
	return &GRPCHandlers{
		chatService: chatService,
	}
}

// Create handles the RPC call to create a new chat.
// It takes a CreateRequest, creates a chat, and returns a CreateResponse with the new chat ID.
func (h *GRPCHandlers) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("rpc Create, request: %+v", req)

	resp, err := h.chatService.CreateChat(ctx, converter.ConvertCreateRequestFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return converter.ConvertCreateChatResponseFromServiceToHandler(resp), nil
}

// Delete handles the RPC call to delete an existing chat.
// It takes a DeleteRequest, deletes the chat, and returns an empty response.
func (h *GRPCHandlers) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("rpc Delete, request: %+v", req)

	err := h.chatService.DeleteChat(ctx, converter.ConvertDeleteRequestFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SendMessage handles the RPC call to send a message to a chat.
// It takes a SendMessageRequest, sends the message, and returns an empty response.
func (h *GRPCHandlers) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("rpc SendMessage, request: %+v", req)

	err := h.chatService.SendMessage(ctx, converter.ConvertSendMessageRequestFromHandlerToService(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
