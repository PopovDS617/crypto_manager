-- +goose Up
create table
    btc (
        id serial primary key,
        price float not null,
        created_at timestamp not null
    );

-- +goose Down
drop table btc;