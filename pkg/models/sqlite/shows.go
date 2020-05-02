package sqlite

import (
    "database/sql"

    "curtaincall.tech/pkg/models"
)

type ShowModel struct {
    DB *sql.DB
}

func (m *ShowModel) Latest() ([]*models.Show, error) {
    stmt := `SELECT show_id, name, company FROM shows
             ORDER BY show_id ASC LIMIT 4`

    rows, err := m.DB.Query(stmt)
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
