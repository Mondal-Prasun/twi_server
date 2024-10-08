-- name: CreateUser :one
INSERT INTO users (id,username,password,image,email,createdAt,updatedAt,accessToken)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id,accessToken;


-- name: UploadUserImage :one
UPDATE users
SET image = $2
WHERE id = $1
RETURNING id, username, image, accessToken;



-- name: GetUserDetails :one
SELECT id,username,image FROM users WHERE id = $1;


-- name: GetUserPasswordByEmail :one
SELECT password,id FROM users WHERE email = $1;

-- name: RefreshUserAccessToken :one
UPDATE users
SET accessToken = $2
WHERE id = $1
RETURNING id,username,email,image, accessToken;
