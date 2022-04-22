package auth

import "fmt"

type Service struct {
	repo *Repo
}

func NewService() *Service {
	return &Service{
		repo: NewRepo(),
	}
}

func (s *Service) UpdateUserInfo(phone string, name string) error {
	err := s.repo.UpdateUserInfo(phone, name)
	if err != nil {
		fmt.Println("error updating user info")
		return ErrQueryDB
	}
	return nil
}
