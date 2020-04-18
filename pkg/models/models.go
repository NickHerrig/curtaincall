package models 

import (
    "errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Theater struct {
    ID      int
    Name    string
}

type Show struct {
    ID      int
    Name    string
    Company string
}

type Song struct {
    ID          int
    Name        string
    Number      int
    ActNumber   int
}

type Performer struct {
    ID          int
    Name        string
    Bio         string
}

type Role struct {
    ID          int
    Name        string
    Order       int
}
