package sqlite

import (
    "database/sql"

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
    return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
    return nil, nil
}
