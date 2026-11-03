-- +goose Up
create table members
(
    id          UUID primary key ,
    first_name  varchar(100) not null,
    last_name   varchar(100) not null,
    birthdate   date         not null,
    email       varchar(255) not null unique,
    phone_number varchar(25) not null unique,
    address     text         not null,
    created_by  UUID         not null references users (id) on delete cascade,
    created_at  timestamp not null default current_timestamp,
    updated_at  timestamp not null default current_timestamp
);

-- +goose Down
drop table members;
