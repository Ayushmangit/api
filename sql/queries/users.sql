-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at, updated_at,api_key)
VALUES ($1,$2, $3, $4, $5,encode(sha256(random()::text::bytea),'hex')
)
RETURNING id, name, email, created_at, updated_at,api_key;

-- name: GetUserByAPIKey :one
select id,name,email,created_at,updated_at,api_key from users where api_key = $1;

-- name: DestroyUser :one
delete from users where id = $1 returning id;

-- name: GetAllUsers :many
SELECT id,name,email,created_at,updated_at,api_key
FROM users
ORDER BY created_at desc;

-- name: UpdateUser :one
UPDATE users
SET
		name = COALESCE($2,name),
		email = COALESCE($3,email),
		updated_at = NOW()
WHERE id = $1
returning id,name,email,created_at,updated_at,api_key;
