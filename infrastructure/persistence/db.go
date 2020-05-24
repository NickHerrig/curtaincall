package persistence

import (
    "fmt"

    "curtaincall.tech/domain/entity"
    "curtaincall.tech/domain/repository"

    _ "github.com/mattn/go-sqlite3"
)

type Repositories struct {
    Theater repository.TheaterRepository
    db      *gorm.DB
}

func NewRepositories(DbDriver, DbDsn) (*Repositories, error) {
    db, err := gorm.Open(DbDriver, DbDsn)
    if err != nil {
        return nil, err
    }

    db.LogMode(true)

    return &Repositories{
        Theater: NewTheaterRepository(db),
        db:      db,
    }, nil
}

func (s *Repositories) Close() error {
    return s.db.Close()
}

func (s *Repositories) Automigrate() error {
    return s.db.AutoMigrate(&entity.Theater{}).Error
}
