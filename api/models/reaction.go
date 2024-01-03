package models

import "time"

type Reaction struct {
	ID        string
	AuthorID  string
	EntriesID string

	UserName  string
	FirstName string
	LastName  string
	Avatar    string

	Action string

	CreationDate time.Time
}
