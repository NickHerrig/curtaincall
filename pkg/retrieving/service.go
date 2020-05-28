package retrieving

type Service interface {
    RetrieveAllTheaters() ([]Theater, error)
}

type Repository interface {
    RetrieveAllTheaters() ([]Theater, error)
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) RetrieveAllTheater() ([]Theater, error) {
     return s.r.RetrieveAllTheater()
}
