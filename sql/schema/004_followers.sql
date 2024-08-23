-- +goose Up
CREATE TABLE user_follows (
    followerId UUID NOT NULL,
    followedId UUID NOT NULL,
    followDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (followerId, followedId),
    FOREIGN KEY (followerId) REFERENCES users(id),
    FOREIGN KEY (followedId) REFERENCES users(id)
);


-- +goose Down
DROP TABLE user_follow;