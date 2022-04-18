package cinema

import (
    "fmt"
    "github.com/templateOfService/models"
)

type Service struct {
    repo *Repo
}

func NewService() *Service {
    return &Service{
        repo: NewRepo(),
    }
}

func (s *Service) ListCinemas(filmId int) ([]models.Cinema, error) {
    if filmId > 0 {
        cinemas, err := s.repo.ListCinemasByFilmId(filmId)
        if err != nil {
            return nil, ErrQueryDB
        }
        return cinemas, nil
    }

    cinemas, err := s.repo.ListCinemas()
    if err != nil {
        fmt.Printf("error listing cinemas: %s\n", err.Error())
        return nil, ErrQueryDB
    }

    return cinemas, nil
}

func (s *Service) GetCinemaById(id int) (*models.Cinema, error) {
    cinema, err := s.repo.GetCinemaById(id)
    if err != nil {
        fmt.Printf("error getting cinema by id: %s\n", err.Error())
        return nil, ErrQueryDB
    }
    return cinema, nil
}
