package adding 
import (
    "errors"
)

var ErrDuplicate = errors.New("beer already exists")

type Service interface {
    AddTheater(Theater) error
}

type Repository interface {
    AddTheater(Theater) error
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
