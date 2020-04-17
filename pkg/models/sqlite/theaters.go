package sqlite

import (
    "database/sql"

    "curtaincall.tech/pkg/models"
)

type TheaterModel struct {
    DB *sql.DB
}

func (m *TheaterModel) Insert(name string) (int, error) {
    stmt := `INSERT INTO theaters (name) VALUES (?)`

    result, err := m.DB.Exec(stmt, name)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return int(id), nil
}

func (m *TheaterModel) Get(id int) (*models.Theater, error) {
    return nil, nil
}
