package log

const (
	queryCreateAPILog = `
		INSERT INTO chats.api_chat_log
			(action_type, request_data, response_data)
		VALUES
			($1, $2, $3);
	`
)
