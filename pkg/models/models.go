package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: Invalid Credentials Entered")
	ErrDuplicateEmail     = errors.New("models: Email already exists")
)

type Theater struct {
	ID   int
	Name string
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassowrd []byte
	Created        time.Time
	Active         bool
}

type Show struct {
	ID      int
	Name    string
	Company string
}
