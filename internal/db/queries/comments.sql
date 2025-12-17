-- name: CreateComment :one
INSERT INTO
    "comments" (content)
VALUES
    ($1)
RETURNING
    id,
    content,
    created_at;

-- name: GetAllComments :many
SELECT
    *
FROM
    "comments"
LIMIT
    100;

-- name: GetCommentById :one
SELECT
    *
FROM
    "comments" c
WHERE
    c.id = $1;
