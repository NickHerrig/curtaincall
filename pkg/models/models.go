package models 

import (
    "errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Theater struct {
    ID      int
    name    string
}

type Show struct {
    ID      int
    name    string
    company string
}

type Song struct {
    ID          int
    name        string
    number      int
    actNumber   int
}
