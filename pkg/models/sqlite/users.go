package sqlite

import (
    "database/sql"
    "errors"

    "curtaincall.tech/pkg/models"

    "github.com/mattn/go-sqlite3"
    "golang.org/x/crypto/bcrypt"
)

type UserModel struct {
    DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return err
    }
    stmt := `INSERT INTO users (name, email, hashed_password, created)
             VALUES (?, ?, ?, datetime("now"))`

    _, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
    if err != nil {
        if SqliteError, ok := err.(sqlite3.Error); ok {
            if SqliteError.Code == sqlite3.ErrConstraint {
                return models.ErrDuplicateEmail
            }
        }
        return err
    }
    return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
    var id int
    var hashedPassword []byte
    stmt := `SELECT user_id, hashed_password FROM users WHERE email = ? AND active = True`
    row := m.DB.QueryRow(stmt, email)
    err := row.Scan(&id, &hashedPassword)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return 0, models.ErrInvalidCredentials
        } else {
            return 0, err
        }
    }

    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if err != nil {
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return 0, models.ErrInvalidCredentials
        } else {
            return 0, err
        }
    }

    return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
    u := &models.User{}
    stmt := `SELECT user_id, name, email, created, active FROM users WHERE user_id = ?`
    err := m.DB.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created, &u.Active)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, models.ErrNoRecord
        } else {
            return nil, err
        }
    }

    return u, nil
}
