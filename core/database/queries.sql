-- name: GetUsers :many
SELECT username FROM users;

-- name: GetUserByName :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    ?, ?, ?
)
RETURNING *;