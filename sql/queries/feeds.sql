-- name: AddUserFeed :one
INSERT INTO feeds (id, created_at, updated_at, user_id, name, url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeed :many
SELECT * FROM feeds;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2;