-- +goose Up
create table products
(
    id                  UUID primary key,
    product_name        varchar(255) not null,
    product_description text         not null,
    product_image       text         not null,
    quantity            int          not null,
    price               numeric      not null,
    added_by            UUID         not null references users (id) on delete cascade,
    created_at          timestamp not null default current_timestamp,
    updated_at          timestamp not null default current_timestamp
);

-- +goose Down
drop table products;
