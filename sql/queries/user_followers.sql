-- name: FollowUser :exec
INSERT INTO user_follows (followerId, followedId)
VALUES ($1, $2);


-- name: UnFollowUser :exec
DELETE FROM user_follows
WHERE followerId = $1 AND followedId = $2;


-- name: FollowedIds :many
SELECT followedId
FROM user_follows
WHERE followerId = $1;

-- name: FollowerIds :many
SELECT followerId
FROM user_follows
WHERE followedId = $1;

-- name: FollowerCount :one
SELECT COUNT(followerId) AS followerCount
FROM user_follows
WHERE followedId = $1;
