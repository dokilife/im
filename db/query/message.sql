-- name: CreateMessage :one
INSERT INTO messages (
  "from",
  message
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM messages ORDER BY created_at DESC LIMIT 1;