-- name: GetAllUsers :many
select id, username, role, created_at, updated_at
from users;

-- name: GetUserByUsername :one
select id, username, role, created_at, updated_at
from users
where username = $1;
