package sqlite

import (
    "database/sql"

    "curtaincall.tech/pkg/models"
)

type UserModel struct {
    DB *sql.DB
}


func (m *UserModel) Insert(name string) error {
    return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
    return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
    return nil, nil
}
