package adding 

type Service interface {
    AddTheater(Theater) error
    AddShow(Show)       error
}

type Repository interface {
    AddTheater(Theater) error
    AddShow(Show)       error
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) AddTheater(t Theater) error {
    err := s.r.AddTheater(t)
    if err != nil {
        return err
    }
     return nil
}

func (s *service) AddShow(sh Show) error {
    err := s.r.AddShow(sh)
    if err != nil {
        return err
    }
     return nil
}
