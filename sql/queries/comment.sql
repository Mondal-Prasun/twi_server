-- name: GiveComment :one
INSERT INTO comments (id,comment,postId,userId,createdAt,updatedAt)
VALUES ($1, $2, $3, $4, $5, $6 )
RETURNING *;

-- name: GetAllComment :many
SELECT *
FROM comments
WHERE postId = $1;
