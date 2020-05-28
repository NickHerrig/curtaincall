package sqlite

import (
    "database/sql"

    "curtaincall.tech/pkg/creating"
    "curtaincall.tech/pkg/retrieving"

    _ "github.com/mattn/go-sqlite3"
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

func (s *Storage) CreateTheater(t creating.Theater) (int, error) {
    stmt := `INSERT INTO theaters (name, address, description)
             VALUES (?, ?, ?)`

    result, err := s.db.Exec(stmt, t.Name, t.Address, t.Description)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

	return int(id), nil
}

func (s *Storage) CreateShow(sh creating.Show) (int, error) {
    stmt := `INSERT INTO shows (name, company, description)
             VALUES (?, ?, ?)`

    result, err := s.db.Exec(stmt, sh.Name, sh.Company, sh.Description)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

	return int(id), nil
}


func (s *Storage) RetrieveAllTheaters() ([]*retrieving.Theater, error) {
    stmt := `SELECT theater_id, name, address, description FROM theaters`

    rows, err := s.db.Query(stmt)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    theaters := []*retrieving.Theater{}

    for rows.Next() {
        t := &retrieving.Theater{}
        err = rows.Scan(&t.ID, &t.Name, &t.Address, &t.Description)
        if err != nil {
            return nil, err
        }
        theaters = append(theaters, t)
    }

    if err = rows.Err(); err!= nil {
        return nil, err
    }

    return theaters, nil    
}
