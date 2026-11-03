-- +goose Up
create table refresh_tokens
(
    token      text primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    user_id    UUID      not null references users (id) on delete cascade,
    expires_at timestamp,
    revoked_at timestamp
);

-- +goose Down
drop table refresh_tokens;
