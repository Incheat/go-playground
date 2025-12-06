-- name: CreateUser :exec
INSERT INTO users (id, email, password_hash, name)
VALUES (?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT id, email, password_hash, name, created_at
FROM users
WHERE email = ?;
