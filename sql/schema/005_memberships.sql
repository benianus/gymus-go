-- +goose Up
create table memberships
(
    id                 UUID primary key,
    member_id          UUID not null unique references members (id) on delete cascade,
    start_date         date         default current_date,
    end_date           date         default (current_date + 30),
    membership_type_id UUID not null references membership_types (id) on delete cascade,
    created_by         UUID not null references users (id) on delete cascade,
    created_at         timestamp not null default current_timestamp,
    updated_at         timestamp not null default current_timestamp
);

-- +goose Down
drop table memberships;
