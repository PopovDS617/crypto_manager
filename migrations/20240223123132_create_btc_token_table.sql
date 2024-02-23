-- +goose Up
create table
    btcusdt (
        id serial primary key,
        price float not null,
        created_at timestamp not null
    );

-- +goose Down
drop table btcusdt;