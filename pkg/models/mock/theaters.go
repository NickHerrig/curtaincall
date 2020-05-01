package mock

import (
    "curtaincall.tech/pkg/models"
)

var mockTheater = &models.Theater{
    ID:    1,
    Name: "Des Moines Civic Center",
}

type TheaterModel struct{}

func (m *TheaterModel) Insert(name string) (int, error) {
    return 2, nil
}

func (m *TheaterModel) Get(id int) (*models.Theater, error) {
    switch id {
    case 1:
        return mockTheater, nil
    default:
        return nil, models.ErrNoRecord
    }
}

func (m *TheaterModel) Latest() ([]*models.Theater, error) {
    return []*models.Theater{mockTheater}, nil
}
