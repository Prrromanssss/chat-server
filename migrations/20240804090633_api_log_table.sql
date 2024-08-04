-- +goose Up
CREATE TABLE chats.api_chat_log (
    id integer GENERATED ALWAYS AS IDENTITY,
    action_type VARCHAR(50) NOT NULL,
    request_data JSONB NOT NULL,
    response_data JSONB,
    timestamp TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE chats.api_chat_log;
