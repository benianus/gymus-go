-- +goose Up
create table membership_types
(
    id         UUID primary key,
    name       varchar(255) not null,
    price      numeric      not null check ( price > 0 ),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- +goose Down
drop table membership_types;
