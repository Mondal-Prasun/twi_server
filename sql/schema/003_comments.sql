-- +goose Up
CREATE TABLE comments (
    id UUID PRIMARY KEY NOT NULL,
    comment TEXT NOT NULL,
    postId UUID NOT NULL,
    userId UUID NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    FOREIGN KEY (postId) REFERENCES posts(id),
    FOREIGN KEY (userId) REFERENCES users(id)
);



-- +goose Down
DROP TABLE comments;