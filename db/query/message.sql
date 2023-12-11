-- name: CreateMessage :one
INSERT INTO messages (
  "from",
  message
) VALUES (
  $1, $2
) RETURNING *;