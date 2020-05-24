package application

import (
    "curtaincall.tech/domain/entity"
    "curtaincall.tech/domain/repository"
)

type theaterApp struct {
    tr repository.TheaterRepository
}

var _ TheaterAppInterface = &theaterApp{}

type TheaterAppInterface interface {
    CreateTheater(*entity.Theater) (*entity.Theater, map[string]string)
    RetriveTheater(*entity.Theater, error)
    RetriveAllTheaters() ([]entity.Theater, error)
    UpdateTheater(*entity.Theater) (*entity.Theater, map[string]string)
    DeleteTheater(int) error
}

func (f *thaterApp) RetriveAllTheaters() ([]entity.Theater, error) {
    return f.tr.RetriveAllTheaters()
{
