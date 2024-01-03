package lib

import (
	"database/sql"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/models"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type Middleware func(Handler) Handler

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
type Credentials struct {
	Identifiers string `json:"identifiers"`
	Password    string `json:"password"`
}
type Payload struct {
	User    models.User
	Session models.Session
}

type DB struct {
	Instance *sql.DB
	Err      error
}
