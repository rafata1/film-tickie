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
