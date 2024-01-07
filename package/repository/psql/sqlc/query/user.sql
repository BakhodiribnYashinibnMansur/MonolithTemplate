-- name: ListUser :many
SELECT * FROM "account";

-- name: GetAccountGyId :one
SELECT * FROM "account" WHERE id = $1;

-- name: CountAccount :one
SELECT count(*) FROM "account";

-- name: CreateAccount :exec
INSERT INTO "account" (full_name, phone_number, birthday_date) VALUES ($1, $2, $3);

-- name: UpdateAccount :exec
UPDATE "account" SET full_name = $2, phone_number = $3, birthday_date = $4 WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM "account" WHERE id = $1;

-- name: SoftDeleteAccount :exec
UPDATE "account" SET deleted_at = now() WHERE id = $1;

-- name: HardDeleteAccount :exec
DELETE FROM "account" WHERE id = $1;

-- name: GetAccountByName :one
SELECT * FROM "account" WHERE phone_number = $1;