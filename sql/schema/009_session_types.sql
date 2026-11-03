-- +goose Up
create table session_types
(
    id         UUID primary key,
    name       varchar(100) not null,
    price      numeric      not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- +goose Down
drop table session_types;
