package models

import "time"

type Post struct {
	ID       string
	Image    any
	Title    string
	Content  string
	AuthorID string

	Likes    []string
	DisLikes []string

	Comments []string
	Category []string

	Username  string
	FirstName string
	LastName  string
	Avatar    string

	LikeStatus string

	CreationDate time.Time
}
