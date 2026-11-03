-- +goose Up
create table users
(
    id         UUID          primary key,
    username   varchar(255) not null unique,
    password   text not null,
    role       varchar(50)  not null default 'OWNER',
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- +goose Down
drop table users;
