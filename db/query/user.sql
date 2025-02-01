-- name: CreateUser :one
INSERT INTO "user" (tenant_id, username, password, role)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;