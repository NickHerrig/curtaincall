package sqlite 

import (
    "database/sql"

    "curtaincall.tech/pkg/models"
)

type TheaterModel struct {
    DB *sql.DB
}

func (m *TheaterModel) Insert(name string) (int, error) {
    return 0, nil
}

func (m *TheaterModel) Get(id int) (*models.Theater, error) {
    return nil, nil
}
