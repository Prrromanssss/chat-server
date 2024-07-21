package chat

const (
	queryCreateChat = `
		INSERT INTO chats.chat DEFAULT VALUES
		RETURNING id;
	`

	queryLinkParticipantsToChat = `
		INSERT INTO chats.chat_participants
			(chat_id, user_id)
		VALUES
			(:chat_id, :user_id);
	`

	queryUnlinkParticipantsFromChat = `
		DELETE FROM chats.chat_participants
		WHERE chat_id = $1
	`

	queryDeleteChat = `
		DELETE FROM chats.chat
		WHERE id = $1
	`

	querySendMessage = `
		INSERT INTO chats.messages
			(sender, message_text, sent_at)
		VALUES
			($1, $2, $3);
	`
)
