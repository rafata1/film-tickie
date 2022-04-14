package models

import "time"

type Schedule struct {
    Id        int       `db:"id"`
    FilmId    int       `db:"film_id"`
    CinemaId  int       `db:"cinema_id"`
    FromTime  time.Time `db:"from_time"`
    ToTime    time.Time `db:"to_time"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}
