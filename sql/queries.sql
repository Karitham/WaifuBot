-- name: GetChars :many
SELECT *
FROM characters
WHERE characters.user_id = $1;

-- name: GetChar :one
SELECT *
FROM characters
WHERE id = $1
    AND characters.user_id = $2;

-- name: InsertChar :exec
INSERT INTO characters ("id", "user_id", "image", "name", "type")
VALUES ($1, $2, $3, $4, $5);

-- name: GiveChar :one
UPDATE characters
SET "type" = 'TRADE',
    "user_id" = @given
WHERE characters.id = @id
    AND characters.user_id = @giver
RETURNING *;

-- name: CreateUser :exec
INSERT INTO users (user_id)
VALUES ($1);