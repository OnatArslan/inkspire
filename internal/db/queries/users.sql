-- name: CreateUser :exec
INSERT INTO
    users (email, password)
VALUES
    ($1, $2);

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    password,
    created_at
FROM
    users
WHERE
    email = $1;

-- name: GetUserByID :one
SELECT
    id,
    email,
    password,
    created_at
FROM
    users
WHERE
    id = $1;
