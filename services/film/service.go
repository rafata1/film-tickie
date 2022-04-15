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

func (s *Service) GetFilmById(id int) (*models.Film, error) {
    film, err := s.repo.GetFilmById(id)
    if err != nil {
        return nil, ErrQueryDB
    }
    return film, nil
}

func (s *Service) ListFilms() ([]models.Film, error) {
    films, err := s.repo.ListFilms()
    if err != nil {
        return nil, ErrQueryDB
    }
    return films, nil
}

func (s *Service) ListFilmsByCategory(category string) ([]models.Film, error) {
    films, err := s.repo.ListFilmsByCategory(category)
    if err != nil {
        return nil, ErrQueryDB
    }
    return films, nil
}
