-- +goose Up
ALTER TABLE chats.chat_participants
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES chats.users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE chats.chat_participants
DROP CONSTRAINT fk_user_id;
