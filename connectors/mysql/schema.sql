create table films
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        not null,
    description text                                not null,
    length      int                                 not null comment 'in seconds',
    opening_day datetime                            not null,
    category    varchar(255)                        not null,
    image_urls  text                                not null,
    created_at  timestamp default CURRENT_TIMESTAMP null,
    updated_at  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint films_name_uindex
        unique (name)
);
