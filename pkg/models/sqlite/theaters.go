package sqlite

import (
    "database/sql"
    "errors"

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
    stmt := `SELECT theater_id, name FROM theaters WHERE theater_id = ?`

    row := m.DB.QueryRow(stmt, id)
    s := &models.Theater{}

    err := row.Scan(&s.ID, &s.Name)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, models.ErrNoRecord
        } else {
            return nil, err
        }
    }
    return s, nil
}
