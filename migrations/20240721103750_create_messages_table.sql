-- +goose Up
CREATE TABLE chats.messages (
    id integer GENERATED ALWAYS AS IDENTITY,
    sender text NOT NULL,
    message_text text NOT NULL,
    sent_at timestamp DEFAULT now(),

    PRIMARY KEY(id)
);


-- +goose Down
DROP TABLE chats.messages;
