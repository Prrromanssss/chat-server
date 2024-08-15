package converter

import (
	"github.com/Prrromanssss/chat-server/internal/model"
	modelRepo "github.com/Prrromanssss/chat-server/internal/repository/chat/model"
)

// ConvertCreateChatResponseFromRepoToService converts a CreateChatResponse from the repository layer
// to a CreateChatResponse used in the service layer.
func ConvertCreateChatResponseFromRepoToService(params modelRepo.CreateChatResponse) model.CreateChatResponse {
	return model.CreateChatResponse{
		ChatID: params.ChatID,
	}
}

// ConvertCreateUsersForChatParamsFromServiceToRepo converts CreateUsersForChatParams
// from the service layer format to the repository layer format.
func ConvertCreateUsersForChatParamsFromServiceToRepo(params model.CreateUsersForChatParams) modelRepo.CreateUsersForChatParams {
	return modelRepo.CreateUsersForChatParams{
		Emails: params.Emails,
	}
}

// ConvertSendMessageParamsFromServiceToRepo converts SendMessageParams
// from the service layer format to the repository layer format.
func ConvertSendMessageParamsFromServiceToRepo(params model.SendMessageParams) modelRepo.SendMessageParams {
	return modelRepo.SendMessageParams{
		From:   params.From,
		Text:   params.Text,
		SentAt: params.SentAt,
	}
}

// ConvertCreateUsersForChatResponseFromRepoToService converts CreateUsersForChatResponse
// from the repository layer format to the service layer format.
func ConvertCreateUsersForChatResponseFromRepoToService(params modelRepo.CreateUsersForChatResponse) model.CreateUsersForChatResponse {
	return model.CreateUsersForChatResponse{
		UserIDs: params.UserIDs,
	}
}

// ConvertLinkParticipantsToChatParamsFromServiceToRepo converts LinkParticipantsToChatParams
// from the service layer format to the repository layer format.
func ConvertLinkParticipantsToChatParamsFromServiceToRepo(params model.LinkParticipantsToChatParams) modelRepo.LinkParticipantsToChatParams {
	return modelRepo.LinkParticipantsToChatParams{
		ChatID:  params.ChatID,
		UserIDs: params.UserIDs,
	}
}

// ConvertUnlinkParticipantsFromChatParamsFromServiceToTepo converts UnlinkParticipantsFromChatParams
// from the service layer format to the repository layer format.
func ConvertUnlinkParticipantsFromChatParamsFromServiceToRepo(params model.UnlinkParticipantsFromChatParams) modelRepo.UnlinkParticipantsFromChatParams {
	return modelRepo.UnlinkParticipantsFromChatParams{
		ChatID: params.ChatID,
	}
}

// ConvertDeleteChatParamsFromServiceToRepo converts DeleteChatParams
// from the service layer format to the repository layer format.
func ConvertDeleteChatParamsFromServiceToRepo(params model.DeleteChatParams) modelRepo.DeleteChatParams {
	return modelRepo.DeleteChatParams{
		ChatID: params.ChatID,
	}
}
