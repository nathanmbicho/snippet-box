package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Snippet table db columns
type Snippet struct {
	ID               int
	Title, Content   string
	Created, Expires time.Time
}
