-- +goose Up
-- +goose StatementBegin
create table gotes (
   id bigserial not null primary key,
   user_id bigint not null ,
   name varchar(255) not null ,
   text text not null ,
   created_at bigint,
   updated_at bigint,

   foreign key (user_id) references users (id) on DELETE CASCADE

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table gotes;
-- +goose StatementEnd
