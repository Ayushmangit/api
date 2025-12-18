-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at, updated_at)
VALUES ($1,$2, $3, $4, $5)
RETURNING id, name, email, created_at, updated_at;

-- name: DestroyUser :one
delete from users where id = $1 returning id;

-- name: GetAllUsers :many
SELECT id,name,email,created_at,updated_at
FROM users
ORDER BY created_at desc;

-- name: UpdateUser :one
UPDATE users
SET
		name = COALESCE($2,name),
		email = COALESCE($3,email),
		updated_at = NOW()
WHERE id = $1
returning id,name,email,created_at,updated_at;
