package schedule

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/templateOfService/connectors/mysql"
    "github.com/templateOfService/models"
)

type Repo struct {
    conn *sqlx.DB
}

func NewRepo() *Repo {
    return &Repo{
        conn: mysql.GetMySQLInstance(),
    }
}

var listSchedulesByCinemaIdAndFilmIdQuery = "SELECT * FROM schedules WHERE cinema_id = ? AND film_id = ?"

func (r *Repo) ListSchedulesByCinemaIdAndFilmId(cinemaId int, filmId int) ([]models.Schedule, error) {
    var res []models.Schedule
    err := r.conn.Select(&res, listSchedulesByCinemaIdAndFilmIdQuery, cinemaId, filmId)
    return res, err
}

var listSchedulesByCinemaIdQuery = "SELECT * FROM schedules WHERE cinema_id = ?"

func (r *Repo) ListSchedulesByCinemaId(cinemaId int) ([]models.Schedule, error) {
    var res []models.Schedule
    err := r.conn.Select(&res, listSchedulesByCinemaIdQuery, cinemaId)
    return res, err
}

var listSchedulesByFilmIdQuery = "SELECT * FROM schedules WHERE film_id = ?"

func (r *Repo) ListSchedulesByFilmId(filmId int) ([]models.Schedule, error) {
    var res []models.Schedule
    err := r.conn.Select(&res, listSchedulesByFilmIdQuery, filmId)
    if err != nil {
        fmt.Println(err)
    }
    return res, err
}
