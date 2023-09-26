-- name: AddUserFeed :one
INSERT INTO feeds (id, created_at, updated_at, user_id, name, url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeed :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;