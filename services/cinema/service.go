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

func (s *Service) ListCinemas() ([]models.Cinema, error) {
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
