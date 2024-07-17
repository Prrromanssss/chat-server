package grpc

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

type GRPCHandlers struct {
	pb.UnimplementedChatV1Server
}

func NewGRPCHandlers() pb.ChatV1Server {
	return &GRPCHandlers{}
}

func (h *GRPCHandlers) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("rpc Create, request: %+v", req)

	return &pb.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (h *GRPCHandlers) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("rpc Delete, request: %+v", req)

	return &emptypb.Empty{}, nil
}

func (h *GRPCHandlers) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("rpc SendMessage, request: %+v", req)

	return &emptypb.Empty{}, nil
}
