package creating

type Service interface {
    CreateTheater(Theater) (int, error)
    CreateShow(Show)       (int, error)
}

type Repository interface {
    CreateTheater(Theater) (int, error)
    CreateShow(Show)       (int, error)
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) CreateTheater(t Theater) (int, error) {
    id, err := s.r.CreateTheater(t)
    if err != nil {
        return 0, err
    }
     return int(id), nil
}

func (s *service) CreateShow(sh Show) (int, error) {
    id, err := s.r.CreateShow(sh)
    if err != nil {
        return 0, err
    }
     return int(id), nil
}
