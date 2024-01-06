package models

import "time"

type Reaction struct {
	ID        string
	AuthorID  string
	EntriesID string

	User User

	Action string

	CreationDate time.Time
}
