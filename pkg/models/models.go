package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
)

type Theater struct {
	ID   int
	Name string
}

type Show struct {
	ID      int
	Name    string
	Company string
}
