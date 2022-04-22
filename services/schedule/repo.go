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

var listAllSchedulesQuery = "SELECT * FROM schedules"

func (r *Repo) ListAllSchedules() ([]models.Schedule, error) {
	var res []models.Schedule
	err := r.conn.Select(&res, listAllSchedulesQuery)
	return res, err
}

var getOrdersByScheduleIdQuery = "SELECT * FROM orders WHERE schedule_id = ?"

func (r *Repo) GetOrdersByScheduleId(scheduleId int) ([]models.Order, error) {
	var res []models.Order
	err := r.conn.Select(&res, getOrdersByScheduleIdQuery, scheduleId)
	return res, err
}

var createOrdersQuery = `INSERT INTO orders (phone_number, schedule_id, seat_code, status) VALUES (:phone_number, 
:schedule_id, :seat_code, :status)`

func (r *Repo) CreateOrders(orders []models.Order) error {
	_, err := r.conn.NamedExec(createOrdersQuery, orders)
	return err
}

var expiredUserPreviousOrdersQuery = "UPDATE orders SET status = ? WHERE phone_number = ? AND schedule_id = ?"

func (r *Repo) ExpireUserPreviousOrders(phone string, scheduleId int) error {
	_, err := r.conn.Exec(expiredUserPreviousOrdersQuery, models.Expired, phone, scheduleId)
	return err
}

var approveOrdersQuery = "UPDATE orders SET status = ? WHERE phone_number = ? AND schedule_id = ? AND status = ?"

func (r *Repo) ApproveOrders(phone string, scheduleId int) error {
	_, err := r.conn.Exec(approveOrdersQuery, models.Approved, phone, scheduleId, models.Holding)
	return err
}
