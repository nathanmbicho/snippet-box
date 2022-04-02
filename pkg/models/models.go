package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

// Snippet table db columns
type Snippet struct {
	ID               int
	Title, Content   string
	Created, Expires time.Time
}

//User type table db columns
type User struct {
	ID             int
	Name, Email    string
	HashedPassword []byte
	Created        time.Time
}
