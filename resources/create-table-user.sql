create table users
(
    id      serial
        primary key,
    login    varchar(200) not null,
    password varchar(200)  not null,
    customer varchar(200) not null
);