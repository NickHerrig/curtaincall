package sqlite

import (
    "database/sql"
    "errors"

    "curtaincall.tech/pkg/models"
)

type ShowModel struct {
    DB *sql.DB
}

func (m *ShowModel) Latest(id int) ([]*models.Show, error) {
    stmt := `SELECT shows.show_id, shows.name, shows.company 
             FROM theaters JOIN theaters_shows_bridge USING ( theater_id )
             JOIN shows USING ( show_id )
             WHERE theaters.theater_id = ?`

    rows, err := m.DB.Query(stmt, id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    shows := []*models.Show{}

    for rows.Next() {
        s := &models.Show{}
        err = rows.Scan(&s.ID, &s.Name, &s.Company)
        if err != nil {
            return nil, err
        }
        shows = append(shows, s)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return shows, nil
}

func (m *ShowModel) Get(id int) (*models.Show, error) {
    stmt := `SELECT show_id, name, company FROM shows WHERE show_id = ?`

    row := m.DB.QueryRow(stmt, id)
    s := &models.Show{}

    err := row.Scan(&s.ID, &s.Name, &s.Company)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, models.ErrNoRecord
        } else {
            return nil, err
        }
    }
    return s, nil
}
