package models

import "time"

type Comments struct {
	ID       string
	Content  string
	AuthorID string
	PostID   string

	Likes    []string
	DisLikes []string

	Reply []string

	User User

	CreationDate time.Time
}
