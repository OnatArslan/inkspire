-- name: CreateComment :one
INSERT INTO
    comments (content)
VALUES
    ($1)
RETURNING
    id,
    content,
    created_at;
