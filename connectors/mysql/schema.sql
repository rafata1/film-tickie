create table cinemas
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        not null,
    description text                                not null,
    address     text                                not null,
    created_at  timestamp default CURRENT_TIMESTAMP null,
    updated_at  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
)
    auto_increment = 24;

create table films
(
    id          int auto_increment
        primary key,
    name        varchar(255)                            not null,
    description text                                    not null,
    length      int                                     not null comment 'in seconds',
    opening_day timestamp default '0000-00-00 00:00:00' not null on update CURRENT_TIMESTAMP,
    category    varchar(255)                            not null,
    image_urls  text                                    not null,
    created_at  timestamp default CURRENT_TIMESTAMP     null,
    updated_at  timestamp default CURRENT_TIMESTAMP     null on update CURRENT_TIMESTAMP,
    constraint films_name_uindex
        unique (name)
)
    auto_increment = 54;

create index films_category_index
    on films (category);

create table schedules
(
    id         int auto_increment
        primary key,
    cinema_id  int                                 not null,
    film_id    int                                 not null,
    from_time  timestamp                           not null on update CURRENT_TIMESTAMP,
    to_time    timestamp                           null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp                           null on update CURRENT_TIMESTAMP
)
    auto_increment = 24;

create index schedules__film_id_index
    on schedules (film_id);

create index schedules_cinema_film_id_index
    on schedules (cinema_id, film_id);