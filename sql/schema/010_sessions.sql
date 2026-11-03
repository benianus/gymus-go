-- +goose Up
create table sessions
(
    id              UUID primary key,
    full_name       varchar(255) not null,
    session_type_id UUID          not null references session_types (id) on delete cascade,
    created_by      UUID          not null references users (id) on delete cascade,
    created_at      timestamp not null default current_timestamp,
    updated_at      timestamp not null default current_timestamp
);

-- +goose Down
drop table sessions;
