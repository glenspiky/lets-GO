package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching models found")

type Snippet struct {
	ID      int
	Title   string
	content string
	created time.Time
	Expires time.Time
}
