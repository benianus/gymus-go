-- +goose Up
create table sales
(
    id          UUID primary key,
    product_id  UUID     not null references products (id) on delete cascade,
    quantity    int     not null,
    total_price numeric not null check ( total_price > 0 ),
    made_by     UUID     not null default null references users (id) on delete cascade,
    created_at  timestamp not null default current_timestamp,
    updated_at  timestamp not null default current_timestamp
);

-- +goose Down
drop table sales;
