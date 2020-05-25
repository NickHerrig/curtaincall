package sqlite

import (
    "curtaincall.tech/pkg/adding"

    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type Storage struct {
    db *sql.DB
}

func NewStorage() (*Storage, error) {
    s := new(Storage)
    var err error

	s.db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	if err = s.db.Ping(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) AddTheater(t adding.Theater) error {
    stmt := `INSERT INTO theaters (name, address, description)
             VALUES (?, ?, ?)`

    _, err := s.db.Exec(stmt, t.Name, t.Address, t.Description)
    if err != nil {
        return err
    }

	return nil
}
