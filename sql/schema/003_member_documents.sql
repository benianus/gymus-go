-- +goose Up
create table member_documents
(
    id                     UUID primary key,
    birth_certificate      text not null,
    medical_certificate    text not null,
    personal_photo         text not null,
    parental_authorization text default null,
    member_id              UUID  not null unique references members (id) on delete cascade,
    added_by               UUID  not null references users (id) on delete cascade,
    created_at             timestamp not null default current_timestamp,
    updated_at             timestamp not null default current_timestamp
);

-- +goose Down
drop table member_documents;
