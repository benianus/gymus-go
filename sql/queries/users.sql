-- name: GetAllUsers :many
select id, username, role, created_at, updated_at
from users;

-- name: GetUserById :one
select id, username, password, role, created_at, updated_at
from users;

-- name: GetUserByUsername :one
select id, username, password, role, created_at, updated_at
from users
where username = $1;

-- name: CreateUser :one
insert into users (id, username, password, role, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;
