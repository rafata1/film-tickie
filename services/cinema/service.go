package cinema

type Service struct {
    repo *Repo
}

func NewService() *Service {
    return &Service{
        repo: NewRepo(),
    }
}
