package schedule

import "github.com/templateOfService/models"

type Service struct {
    repo *Repo
}

func NewService() *Service {
    return &Service{
        repo: NewRepo(),
    }
}

func (s *Service) ListSchedules(cinemaId int, filmId int) ([]models.Schedule, error) {
    if cinemaId > 0 && filmId > 0 {
        schedules, err := s.repo.ListSchedulesByCinemaIdAndFilmId(cinemaId, filmId)
        if err != nil {
            return nil, ErrQueryDB
        }
        return schedules, nil
    }

    if cinemaId > 0 {
        schedules, err := s.repo.ListSchedulesByCinemaId(cinemaId)
        if err != nil {
            return nil, ErrQueryDB
        }
        return schedules, nil
    }

    if filmId > 0 {
        schedules, err := s.repo.ListSchedulesByFilmId(filmId)
        if err != nil {
            return nil, ErrQueryDB
        }
        return schedules, nil
    }

    return nil, nil
}