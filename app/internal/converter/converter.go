package converter

import (
	"github.com/Prrromanssss/chat-server/internal/model"
	pb "github.com/Prrromanssss/chat-server/pkg/chat_v1"
)

// ConvertCreateRequestFromHandlerToService converts a CreateRequest from the api layer to a CreateChatParams for the service layer.
func ConvertCreateRequestFromHandlerToService(params *pb.CreateRequest) model.CreateChatParams {
	return model.CreateChatParams{
		Emails: params.Emails,
	}
}

// ConvertCreateChatResponseFromServiceToHandler converts a CreateChatResponse from the service layer to a CreateResponse for the api layer.
func ConvertCreateChatResponseFromServiceToHandler(params model.CreateChatResponse) *pb.CreateResponse {
	return &pb.CreateResponse{
		Id: params.ChatID,
	}
}

// ConvertDeleteRequestFromHandlerToService converts a DeleteRequest from the api layer to a DeleteChatParams for the service layer.
func ConvertDeleteRequestFromHandlerToService(params *pb.DeleteRequest) model.DeleteChatParams {
	return model.DeleteChatParams{
		ChatID: params.Id,
	}
}

// ConvertSendMessageRequestFromHandlerToService converts a SendMessageRequest from the api layer to SendMessageParams for the service layer.
func ConvertSendMessageRequestFromHandlerToService(params *pb.SendMessageRequest) model.SendMessageParams {
	return model.SendMessageParams{
		From:   params.From,
		Text:   params.Text,
		SentAt: params.Timestamp.AsTime(),
	}
}
