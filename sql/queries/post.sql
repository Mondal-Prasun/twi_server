-- name: CreatePost :one
INSERT INTO posts (id,username,userId,contentText,contextImage,likes,createdAt,updatedAt)
VALUES ($1, $2, $3, $4, $5, $6, $7 ,$8)
RETURNING *;
-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1 AND userId = $2;


-- name: LikePost :exec
UPDATE posts
SET likes = likes + 1,
    updatedAt = $2
WHERE id = $1;

