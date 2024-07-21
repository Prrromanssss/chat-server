-- +goose Up
CREATE SCHEMA chats;

CREATE TABLE chats.chat (
    id integer GENERATED ALWAYS AS IDENTITY,
    created_at timestamp DEFAULT now(),

    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE chats.chat;

DROP SCHEMA chats;
