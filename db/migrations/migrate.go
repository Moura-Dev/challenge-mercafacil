package migrations

var Schema = `create table if not exists users
(
    id      serial
        primary key,
    login    varchar(200) unique not null,
    password varchar(200)  not null,
    customer varchar(200) not null
);`
