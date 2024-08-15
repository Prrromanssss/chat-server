package chat

const (
	queryCreateChat = `
		INSERT INTO chats.chat DEFAULT VALUES
		RETURNING id;
	`

	queryCreateUser = `
		WITH ins AS (
			INSERT INTO chats.users 
				(email)
			VALUES 
				($1)
			ON CONFLICT (email) DO NOTHING
			RETURNING id
		)
		SELECT id
		FROM ins
		UNION ALL
		SELECT id
		FROM chats.users
		WHERE email = $1;
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
)
