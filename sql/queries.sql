-- name: GetUserList :many
SELECT *
FROM characters
WHERE characters.user_id = $1;
-- name: GetUserCharsIDs :many
SELECT id
FROM characters
WHERE characters.user_id = $1;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetChar :one
SELECT *
FROM characters
WHERE id = $1
    AND characters.user_id = $2
LIMIT 1;
-- name: InsertChar :exec
INSERT INTO characters ("id", "user_id", "image", "name")
VALUES ($2, $1, $3, $4);
-- name: SetQuote :exec
UPDATE users
SET quote = $2
WHERE users.id = $1;
-- name: SetFavorite :exec
UPDATE users
SET favorite = $2
WHERE users.id = $1;
-- name: SetDate :exec
UPDATE users
SET date = $2
WHERE users.id = $1;
-- name: GetUserProfile :one
SELECT characters.image,
    characters.name,
    users.date,
    users.quote,
    (
        SELECT count(id)
        FROM characters
        WHERE characters.user_id = $1
    ) as count
FROM users
    LEFT JOIN characters ON characters.id = users.favorite
WHERE users.id = $1;
-- name: GiveChar :exec
UPDATE characters
SET user_id = $3
WHERE characters.id = $1
    AND characters.user_id = $2;
-- name: GetDate :one
SELECT users.date
FROM users
WHERE users.id = $1;
-- name: CreateUser :exec
INSERT INTO users (id)
VALUES ($1);