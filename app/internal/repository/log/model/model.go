package model

import "database/sql"

// CreateAPILogParams holds the parameters for logging API actions related to user creation.
type CreateAPILogParams struct {
	Method       string         `db:"action_type"`
	RequestData  string         `db:"request_data"`
	ResponseData sql.NullString `db:"response_data"`
}
