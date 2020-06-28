package sqlite

import (
    "database/sql"
    "errors"

    "curtaincall.tech/pkg/retrieving"

    _ "github.com/mattn/go-sqlite3"
)

var ErrNoRecord = errors.New("Storage: no matching record found")

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

func (s *Storage) RetrieveAllShows() ([]*retrieving.Show, error) {
    stmt := `SELECT show_id, name, company, description FROM shows`

    rows, err := s.db.Query(stmt)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    shows := []*retrieving.Show{}

    for rows.Next() {
        sh := &retrieving.Show{}
        err = rows.Scan(&sh.ID, &sh.Name, &sh.Company, &sh.Description)
        if err != nil {
            return nil, err
        }
        shows = append(shows, sh)
    }

    if err = rows.Err(); err!= nil {
        return nil, err
    }

    return shows, nil    
}
