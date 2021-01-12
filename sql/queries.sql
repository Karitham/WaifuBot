-- name: GetUserList :many
SELECT * FROM Characters WHERE "user_id" = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE "id" = $1
LIMIT 1;

-- name: GetChar :one
SELECT * FROM characters
WHERE "id" = $1 AND "user_id" = $2
LIMIT 1;

-- name: InsertChar :exec
INSERT INTO characters ("id", "user_id", "image", "name")
VALUES ($2, $1, $3, $4);

-- name: SetQuote :exec
INSERT INTO users ("id", "quote")
VALUES ($1, $2)
ON CONFLICT ("id") DO
UPDATE SET "quote" = $2 WHERE users.id = $1;

-- name: SetFavorite :exec
INSERT INTO "users" ("id", "favorite")
VALUES ($1, $2)
ON CONFLICT ("id") DO
UPDATE SET "favorite" = $2 WHERE users.id = $1;

-- name: AddOneToClaimCount :exec
INSERT INTO "users" ("id", "claim_count")
VALUES ($1, 1)
ON CONFLICT ("id") DO
UPDATE SET "claim_count" = users.claim_count + 1 WHERE users.id = $1;

-- name: UpdateUserDate :exec
INSERT INTO "users" ("id", "date")
VALUES ($1, $2)
ON CONFLICT ("id") DO
UPDATE SET "date" = $2 WHERE users.id = $1;

-- name: DeleteChar :exec
DELETE FROM characters
WHERE "id" = $1 AND "user_id" = $2;

-- name: GetUserProfile :many
SELECT a.id, a.user_id, a.image, a.name, users.quote, users.date, users.favorite, users.claim_count FROM
(SELECT id, user_id, image, name FROM characters WHERE "user_id" = $1) AS a
INNER JOIN
users ON users.id = a.user_id;

-- name: GiveChar :exec
UPDATE characters
SET "user_id" = $3
WHERE "id" = $1
AND "user_id" = $2;

-- name: GetDate :one
SELECT "date"
FROM users
WHERE "id" = $1
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users ("id")
VALUES ($1);