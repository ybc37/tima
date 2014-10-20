-- +migrate Up
create table users (
    id              integer         not null    auto_increment,
    username        varchar(50)     not null,
    password_hash   binary(60)      not null,
    first_name      varchar(50)         null,
    last_name       varchar(50)         null,
    email           varchar(254)    not null,
    primary key (id),
    unique index (username)
);

-- +migrate Down
drop table users;
