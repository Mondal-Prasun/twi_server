-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    image TEXT,
    email TEXT NOT NULL UNIQUE,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    accessToken UUID NOT NULL
);

-- +goose Down
DROP TABLE users;