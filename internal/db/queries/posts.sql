-- name: CreatePost :one
INSERT INTO
    posts (title, content)
VALUES
    ($1, $2)
RETURNING
    id,
    title,
    content,
    created_at;
