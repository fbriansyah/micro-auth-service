-- name: CreateUser :one
INSERT INTO users (username, password, fullname)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;
