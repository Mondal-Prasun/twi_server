-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    userId UUID NOT NULL,
    contentText TEXT NOT NULL,
    contextImage TEXT,
    likes INT NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id)
);


-- +goose Down
DROP TABLE posts;