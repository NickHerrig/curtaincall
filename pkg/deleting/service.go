package deleting 

type Service interface {
    DeleteTheater(int) error
}

type Repository interface {
    DeleteTheater(int) error
}

type service struct {
    r Repository
}

func NewService(r Repository) Service {
    return &service{r}
}

func (s *service) DeleteTheater(id int) error {
     return s.r.DeleteTheater(id)
}
