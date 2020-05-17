package sqlite

import (
	"database/sql"
	"errors"

	"curtaincall.tech/pkg/models"
)

type TheaterModel struct {
	DB *sql.DB
}

func (m *TheaterModel) Latest() ([]*models.Theater, error) {
	stmt := `SELECT theater_id, name FROM theaters
             ORDER BY theater_id ASC LIMIT 5`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	theaters := []*models.Theater{}

	for rows.Next() {
		t := &models.Theater{}
		err = rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return nil, err
		}
		theaters = append(theaters, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return theaters, nil
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
	t := &models.Theater{}

	err := row.Scan(&t.ID, &t.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return t, nil
}
