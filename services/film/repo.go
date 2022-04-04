package film

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "github.com/templateOfService/connectors/mysql"
    "github.com/templateOfService/models"
    "log"
)

type Repo struct {
    conn *sqlx.DB
}

func NewRepo() *Repo {
    return &Repo{
        conn: mysql.GetMySQLInstance(),
    }
}

var listFilmsQuery = "SELECT * FROM films"

func (r *Repo) ListFilms() ([]models.Film, error) {
    var res []models.Film
    err := r.conn.Select(&res, listFilmsQuery)
    if err != nil {
        log.Printf("Error querying DB: %s", err.Error())
    }
    return res, err
}
