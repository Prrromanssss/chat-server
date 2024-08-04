package converter

import (
	"database/sql"

	"github.com/Prrromanssss/chat-server/internal/model"
	modelRepo "github.com/Prrromanssss/chat-server/internal/repository/chat/model"
)

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

// ConvertCreateAPILogParamsFromServiceToRepo converts CreateAPILogParams from the service layer
// to the repository layer format.
func ConvertCreateAPILogParamsFromServiceToRepo(params model.CreateAPILogParams) modelRepo.CreateAPILogParams {
	var responseData sql.NullString

	if params.ResponseData != nil {
		responseData = sql.NullString{String: *params.ResponseData, Valid: true}
	} else {
		responseData = sql.NullString{String: "", Valid: false}
	}

	return modelRepo.CreateAPILogParams{
		Method:       params.Method,
		RequestData:  params.RequestData,
		ResponseData: responseData,
	}
}
