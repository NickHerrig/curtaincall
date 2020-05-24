package repository

import (
  "curtaincall.tech/domain/entity"
)

type TheaterRepository interface {
    CreateTheater(*entity.Theater) (*entity.Theater, map[string]string)
    RetriveTheater(*entity.Theater, error)
    RetriveAllTheaters() ([]entity.Theater, error)
    UpdateTheater(*entity.Theater) (*entity.Theater, map[string]string)
    DeleteTheater(int)
}
