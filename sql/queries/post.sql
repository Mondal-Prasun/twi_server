-- name: CreatePost :one
INSERT INTO posts (id,username,userId,contentText,contextImage,likes,createdAt,updatedAt)
VALUES ($1, $2, $3, $4, $5, $6, $7 ,$8)
RETURNING *;