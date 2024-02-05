-- name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email_address = $1;

-- name: UserExistsByEmail :one
SELECT COUNT(*)
FROM users
WHERE email_address = $1;

-- name: InsertUser :exec
INSERT INTO users
    (first_name, last_name, email_address, hash)
VALUES
    ($1, $2, $3, $4);

-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    email_address = $4,
    email_verified = $5,
    hash = $6
WHERE id = $1;
