-- name: getChars :many
SELECT *
FROM characters
WHERE characters.user_id = $1
ORDER BY characters.date DESC
LIMIT $2 OFFSET $3;
-- name: getCharsWhoseIDStartWith :many
SELECT *
FROM characters
WHERE characters.user_id = @user_id
    AND characters.id::varchar LIKE @like_str::string
ORDER BY characters.date DESC
LIMIT @lim OFFSET @off;
-- name: getChar :one
SELECT *
FROM characters
WHERE id = $1
    AND characters.user_id = $2;
-- name: insertChar :exec
INSERT INTO characters ("id", "user_id", "image", "name", "type")
VALUES ($1, $2, $3, $4, $5);
-- name: giveChar :one
UPDATE characters
SET "type" = 'TRADE',
    "user_id" = @given
WHERE characters.id = @id
    AND characters.user_id = @giver
RETURNING *;
-- name: createUser :exec
INSERT INTO users (user_id)
VALUES ($1);
-- name: getUser :one
SELECT *
FROM users
WHERE user_id = $1;
-- name: getProfile :one
SELECT characters.image as favorite_image,
    characters.name as favorite_name,
    characters.id as favorite_id,
    users.date as user_date,
    users.quote as user_quote,
    users.user_id as user_id,
    (
        SELECT count(id)
        FROM characters
        WHERE characters.user_id = $1
    ) as count
FROM users
    LEFT JOIN characters ON characters.id = users.favorite
WHERE users.user_id = $1;