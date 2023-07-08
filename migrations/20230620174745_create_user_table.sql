-- +goose Up
-- +goose StatementBegin
create table users (
    id bigserial not null primary key,
    email varchar(255) unique ,
    password_hash varchar(255) not null ,
    created_at bigint,
    updated_at bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
