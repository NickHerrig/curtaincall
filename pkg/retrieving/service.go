package retrieving

type Service interface {
    RetrieveAllTheaters() ([]*Theater, error)
    RetrieveTheater(int) (*Theater, error)
}

type Repository interface {
    RetrieveAllTheaters() ([]*Theater, error)
    RetrieveTheater(int) (*Theater, error)
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) RetrieveAllTheaters() ([]*Theater, error) {
     return s.r.RetrieveAllTheaters()
}

func (s *service) RetrieveTheater(id int) (*Theater, error) {
     return s.r.RetrieveTheater(id)
}
