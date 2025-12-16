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

-- name: GetAllPosts :many
SELECT
    *
FROM
    posts p;

-- name: GetPostById :one
SELECT
    *
FROM
    posts p
WHERE
    p.id = $1;
