package schedule

import (
	"github.com/templateOfService/models"
	"time"
)

type Service struct {
	repo *Repo
}

func NewService() *Service {
	return &Service{
		repo: NewRepo(),
	}
}

func (s *Service) ListSchedules(cinemaId int, filmId int, onDate *time.Time) ([]models.Schedule, error) {
	schedules, err := s.filterScheduleByCinemaAndFilm(cinemaId, filmId)
	if err != nil {
		return nil, ErrQueryDB
	}
	if onDate != nil {
		schedules = filterSchedulesByDate(schedules, onDate)
	}
	return schedules, nil
}

func filterSchedulesByDate(schedules []models.Schedule, date *time.Time) []models.Schedule {
	var res []models.Schedule
	for _, s := range schedules {
		if isEqualDate(s.FromTime, *date) || isEqualDate(s.ToTime, *date) {
			res = append(res, s)
		}
	}
	return res
}

func isEqualDate(a time.Time, b time.Time) bool {
	yearA, monthA, dayA := a.Date()
	yearB, monthB, dayB := b.Date()
	return yearA == yearB && monthA == monthB && dayA == dayB
}

func (s *Service) filterScheduleByCinemaAndFilm(cinemaId int, filmId int) ([]models.Schedule, error) {
	if cinemaId > 0 && filmId > 0 {
		return s.repo.ListSchedulesByCinemaIdAndFilmId(cinemaId, filmId)
	}

	if cinemaId > 0 {
		return s.repo.ListSchedulesByCinemaId(cinemaId)
	}

	if filmId > 0 {
		return s.repo.ListSchedulesByFilmId(filmId)
	}
	return s.repo.ListAllSchedules()
}

func (s *Service) ListSeats(scheduleId int) ([]int, []int, error) {
	ordersOfSchedule, err := s.repo.GetOrdersByScheduleId(scheduleId)
	if err != nil {
		return nil, nil, err
	}

	notExpiredOrders := removeExpiredOrdersFromList(ordersOfSchedule)

	var holdingSeats []int
	var orderedSeats []int
	for _, o := range notExpiredOrders {
		if o.Status == models.Holding {
			holdingSeats = append(holdingSeats, o.SeatCode)
		}
		if o.Status == models.Approved {
			orderedSeats = append(orderedSeats, o.SeatCode)
		}
	}

	return holdingSeats, orderedSeats, nil
}

func (s *Service) HoldSeats(phone string, scheduleId int, seatCodes []int) error {
	err := s.repo.ExpireUserPreviousOrders(phone, scheduleId)
	if err != nil {
		return err
	}

	ordersOfSchedule, err := s.repo.GetOrdersByScheduleId(scheduleId)
	if err != nil {
		return err
	}

	notExpiredOrders := removeExpiredOrdersFromList(ordersOfSchedule)
	err = validateSeats(seatCodes, notExpiredOrders)
	if err != nil {
		return err
	}

	orders := buildOrders(phone, scheduleId, seatCodes, models.Holding)
	err = s.repo.CreateOrders(orders)
	if err != nil {
		return err
	}
	return nil
}

func removeExpiredOrdersFromList(orders []models.Order) []models.Order {
	var res []models.Order
	for _, o := range orders {
		if o.Status == models.Expired {
			continue
		}
		res = append(res, o)
	}
	return res
}

func validateSeats(seatCodes []int, orders []models.Order) error {
	seatCodeToStatusMap := make(map[int]models.OrderStatus)
	for _, o := range orders {
		seatCodeToStatusMap[o.SeatCode] = o.Status
	}

	for _, sc := range seatCodes {
		if seatCodeToStatusMap[sc] == models.Approved || seatCodeToStatusMap[sc] == models.Holding {
			return ErrSeatsAreNotFree
		}
	}
	return nil
}

func buildOrders(phone string, scheduleId int, seatCodes []int, status models.OrderStatus) []models.Order {
	var res []models.Order
	for _, sc := range seatCodes {
		res = append(res, models.Order{
			Phone:      phone,
			ScheduleId: scheduleId,
			SeatCode:   sc,
			Status:     status,
		})
	}
	return res
}

func (s *Service) ConfirmSeats(phone string, scheduleId int) error {
	err := s.repo.ApproveOrders(phone, scheduleId)
	return err
}

func (s *Service) CancelSeats(phone string, scheduleId int) error {
	err := s.repo.ExpireUserPreviousOrders(phone, scheduleId)
	return err
}
