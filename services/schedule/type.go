package schedule

type BaseRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HoldSeatsRequest struct {
	ScheduleId int   `json:"schedule_id"`
	Seats      []int `json:"seats"`
}

type ConfirmSeatsRequest struct {
	ScheduleId int `json:"schedule_id"`
}

type CancelSeatsRequest struct {
	ScheduleId int `json:"schedule_id"`
}

type ListSeatsResponse struct {
	HoldingSeats []int `json:"holding_seats"`
	OrderedSeats []int `json:"ordered_seats"`
}
