-- name: SaveRefreshToken :one
insert into refresh_tokens (token, created_at, updated_at, user_id, revoked_at, expires_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetRefreshTokenByName :one
select token, created_at, updated_at, user_id, revoked_at, expires_at
from refresh_tokens
where token = $1;

-- name: GetRefreshTokenByUserId :one
select token, created_at, updated_at, user_id, revoked_at, expires_at
from refresh_tokens
where user_id = $1;

-- name: RevokeRefreshToken :exec
update refresh_tokens set revoked_at = current_timestamp
where token = $1;
