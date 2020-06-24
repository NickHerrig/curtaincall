package sqlite

import (
    "database/sql"
    "errors"

    "curtaincall.tech/pkg/creating"
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

    ts := []*retrieving.Theater{}

    for rows.Next() {
        t := &retrieving.Theater{}
        err = rows.Scan(&t.ID, &t.Name, &t.Address, &t.Description)
        if err != nil {
            return nil, err
        }
        ts = append(ts, t)
    }

    if err = rows.Err(); err!= nil {
        return nil, err
    }

    return ts, nil    
}

func (s *Storage) RetrieveTheater(id int) (*retrieving.Theater, error) {
    stmt := `SELECT theater_id, name, address, description FROM theaters
             WHERE theater_id = ?`

    row := s.db.QueryRow(stmt, id)

    t := &retrieving.Theater{}

    err := row.Scan(&t.ID, &t.Name, &t.Address, &t.Description)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }

    return t, nil
}

func (s *Storage) DeleteTheater(id int) error {
    stmt := `DELETE FROM theaters WHERE theater_id = ?`
    result, err := s.db.Exec(stmt, id)
    if err != nil {
        return err
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rows < 1 {
        return ErrNoRecord
    }

    return nil
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

func (s *Storage) RetrieveShow(id int) (*retrieving.Show, error) {
    stmt := `SELECT show_id, name, company, description FROM shows 
             WHERE show_id = ?`

    row := s.db.QueryRow(stmt, id)

    sh := &retrieving.Show{}

    err := row.Scan(&sh.ID, &sh.Name, &sh.Company, &sh.Description)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }

    return sh, nil
}
