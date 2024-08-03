-- +goose Up
CREATE TABLE chats.users (
    id integer GENERATED ALWAYS AS IDENTITY,
    email varchar(255) UNIQUE NOT NULL,
    created_at timestamp DEFAULT now(),

    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE chats.users;
