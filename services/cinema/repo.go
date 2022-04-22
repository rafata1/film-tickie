package cinema

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

var listCinemasByFilmIdQuery = `
    SELECT DISTINCT c.id, c.name, c.description, c.address, c.image_urls, c.created_at, c.updated_at FROM cinemas c
    JOIN schedules s ON c.id = s.cinema_id
    WHERE s.film_id = ? AND s.from_time > NOW();
`

func (r *Repo) ListCinemasByFilmId(filmId int) ([]models.Cinema, error) {
	var res []models.Cinema
	err := r.conn.Select(&res, listCinemasByFilmIdQuery, filmId)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
