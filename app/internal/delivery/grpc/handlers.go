package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Prrromanssss/chat-server/internal/models"
	"github.com/Prrromanssss/chat-server/internal/repository"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

type GRPCHandlers struct {
	pb.UnimplementedChatV1Server
	repo repository.ChatRepository
}

func NewGRPCHandlers(repo repository.ChatRepository) pb.ChatV1Server {
	return &GRPCHandlers{repo: repo}
}

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

func (h *GRPCHandlers) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("rpc Delete, request: %+v", req)

	err := h.repo.DeleteChat(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

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
