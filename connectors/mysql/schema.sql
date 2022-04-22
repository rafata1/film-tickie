create table cinemas
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        not null,
    description text                                not null,
    address     text                                not null,
    image_urls  text                                not null,
    updated_at  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    created_at  timestamp default CURRENT_TIMESTAMP null
)
    auto_increment = 84;

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
    auto_increment = 234;

create index films_category_index
    on films (category);

create table orders
(
    id           int auto_increment
        primary key,
    phone_number varchar(20)                             not null,
    schedule_id  int                                     not null,
    seat_code    int                                     null,
    status       varchar(20)                             not null,
    created_at   timestamp default CURRENT_TIMESTAMP     not null,
    updated_at   timestamp default '0000-00-00 00:00:00' not null on update CURRENT_TIMESTAMP,
    constraint orders_phone_number_schedule_id_seat_code_uindex
        unique (phone_number, schedule_id, seat_code)
);

create index orders_schedule_id_index
    on orders (schedule_id);

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
    auto_increment = 204;

create index schedules__film_id_index
    on schedules (film_id);

create index schedules_cinema_film_id_index
    on schedules (cinema_id, film_id);

create table users
(
    id           int auto_increment
        primary key,
    phone_number varchar(20)                             not null,
    name         varchar(100)                            not null,
    created_at   timestamp default CURRENT_TIMESTAMP     not null,
    updated_at   timestamp default '0000-00-00 00:00:00' not null on update CURRENT_TIMESTAMP,
    constraint users_name_uindex
        unique (name),
    constraint users_phone_number_uindex
        unique (phone_number)
)
    auto_increment = 14;

