package cinema

import (
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

var listCinemasQuery = "SELECT * FROM cinemas"

func (r *Repo) ListCinemas() ([]models.Cinema, error) {
    var res []models.Cinema
    err := r.conn.Select(&res, listCinemasQuery)
    return res, err
}

var getCinemaByIdQuery = "SELECT * FROM cinemas WHERE id = ?"

func (r *Repo) GetCinemaById(id int) (*models.Cinema, error) {
    var res models.Cinema
    err := r.conn.Get(&res, getCinemaByIdQuery, id)
    return &res, err
}
