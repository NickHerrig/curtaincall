package persistence

import (
    "errors"

    "curtaincall.tech/domain/entity"
    "curtaincall.tech/domain/repository"

    "github.com/jinzhu/gorm"
)

type TheaterRepo struct {
    db *gorm.DB
}

func NewTheaterRepository (db *gorm.DB) *TheaterRepo {
    return &TheaterRepo{db}
}

var _ repository.TheaterRepository = &TheaterRepo{}

func (r *TheaterRepo) retriveAllTheaters() ([]entity.Theater, error) {
    var theaters []entity.Theater
    err := r.db.Debug().Limit(100).Find(&theaters).Error
    if err != nil {
        return nil, err
    }

    if gorm.IsRecordNotFoundError(err) {
        return nil, errors.New("Theater is not found!")
    }
    return theaters, nil
}
