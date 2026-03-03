-- name: CreateUser :one
INSERT INTO users(id, username, password, created_at, updated_at) VALUES (
				UUIDV7(),
				$1,
				$2,
				NOW(),
				NOW()
)
RETURNING *;
