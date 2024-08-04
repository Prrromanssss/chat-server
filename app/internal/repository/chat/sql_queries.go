package chat

const (
	queryCreateChat = `
		INSERT INTO chats.chat DEFAULT VALUES
		RETURNING id;
	`

	queryCreateUser = `
		INSERT INTO chats.users
			(email)
		VALUES
			($1)
		ON CONFLICT (email) DO NOTHING
		RETURNING
			id;
	`

	queryLinkParticipantsToChat = `
		INSERT INTO chats.chat_participants
			(chat_id, user_id)
		VALUES
			($1, $2);
	`

	queryDeleteChat = `
	DELETE FROM chats.chat
	WHERE id = $1;
	`

	queryUnlinkParticipantsFromChat = `
		DELETE FROM chats.chat_participants
		WHERE chat_id = $1;
	`

	querySendMessage = `
		INSERT INTO chats.messages
			(sender, message_text, sent_at)
		VALUES
			($1, $2, $3);
	`

	queryCreateAPILog = `
		INSERT INTO users.api_user_log
			(action_type, request_data, response_data)
		VALUES
			($1, $2, $3);
	`
)
