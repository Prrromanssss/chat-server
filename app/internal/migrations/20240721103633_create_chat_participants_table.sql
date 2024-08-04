-- +goose Up
CREATE TABLE chats.chat_participants (
    chat_id integer NOT NULL,
    user_id integer NOT NULL,

    PRIMARY KEY (chat_id, user_id),
    FOREIGN KEY (chat_id) REFERENCES chats.chat(id)
);

-- +goose Down
DROP TABLE chats.chat_participants;
