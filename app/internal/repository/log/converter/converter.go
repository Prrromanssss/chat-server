package converter

import (
	"database/sql"
	"encoding/json"

	"github.com/Prrromanssss/chat-server/internal/model"
	modelRepo "github.com/Prrromanssss/chat-server/internal/repository/log/model"
)

// ConvertCreateAPILogParamsFromServiceToRepo converts CreateAPILogParams from the service layer
// to the repository layer format.
func ConvertCreateAPILogParamsFromServiceToRepo(params model.CreateAPILogParams) (modelRepo.CreateAPILogParams, error) {
	requestDataBytes, err := json.Marshal(params.RequestData)
	if err != nil {
		return modelRepo.CreateAPILogParams{}, err
	}
	requestData := string(requestDataBytes)

	var responseData sql.NullString
	if params.ResponseData != nil {
		responseDataBytes, err := json.Marshal(params.ResponseData)
		if err != nil {
			return modelRepo.CreateAPILogParams{}, err
		}
		responseData.String = string(responseDataBytes)
		responseData.Valid = true
	} else {
		responseData.Valid = false
	}

	return modelRepo.CreateAPILogParams{
		Method:       params.Method,
		RequestData:  requestData,
		ResponseData: responseData,
	}, nil
}
