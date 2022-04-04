package film

import "github.com/templateOfService/models"

type Service struct {
    repo *Repo
}

func NewService() *Service {
    return &Service{
        repo: NewRepo(),
    }
}

func (s *Service) ListFilms() ([]models.Film, error) {
    films, err := s.repo.ListFilms()
    if err != nil {
        return nil, ErrQueryDB
    }
    return films, nil
}
