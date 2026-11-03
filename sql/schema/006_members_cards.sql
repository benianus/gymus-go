-- +goose Up
create table members_cards
(
    id            UUID primary key,
    member_id     UUID unique  not null references members (id) on delete cascade ,
    membership_id UUID unique not null references memberships (id) on delete cascade,
    created_by    UUID not null references users (id) on delete cascade,
    created_at    timestamp not null default current_timestamp,
    updated_at    timestamp not null default current_timestamp
);

-- +goose Down
drop table members_cards;
