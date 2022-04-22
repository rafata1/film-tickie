package models

import "time"

type OrderStatus string

const (
	Holding  OrderStatus = "HOLDING"
	Approved OrderStatus = "APPROVED"
	Expired  OrderStatus = "EXPIRED"
)

type Order struct {
	Id         int         `db:"id"`
	Phone      string      `db:"phone_number"`
	ScheduleId int         `db:"schedule_id"`
	SeatCode   int         `db:"seat_code"`
	Status     OrderStatus `db:"status"`
	CreatedAt  time.Time   `db:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
}
