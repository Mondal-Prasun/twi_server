-- name: CreateUser :one
INSERT INTO users (id,username,password,image,email,createdAt,updatedAt,accessToken)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;



-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;