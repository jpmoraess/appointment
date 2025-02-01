-- name: CreateTenant :one
INSERT INTO tenant (name, slug)
VALUES ($1, $2)
RETURNING *;

-- name: GetTenant :one
SELECT * FROM tenant
WHERE id = $1 LIMIT 1;

-- name: ListTenants :many
SELECT * FROM tenant
ORDER BY id
LIMIT $1
OFFSET $2;
