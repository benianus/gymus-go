-- name: GetAllUsers :many
select id, username, role, created_at, updated_at
from users;
