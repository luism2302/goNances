-- name: CreateUser :one
INSERT INTO users(id, username, hashed_password, created_at, updated_at) VALUES (
				UUIDV7(),
				$1,
				$2,
				NOW(),
				NOW()
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: AssignTokenToUser :exec
UPDATE users SET session_token = $1 WHERE id = $2;
