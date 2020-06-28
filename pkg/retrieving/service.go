package retrieving

type Service interface {
    RetrieveAllShows() ([]*Show, error)
}

type Repository interface {
    RetrieveAllShows() ([]*Show, error)
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) RetrieveAllShows() ([]*Show, error) {
     return s.r.RetrieveAllShows()
}
