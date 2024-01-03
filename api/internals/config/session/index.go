package session

import (
	"time"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errHandler "learn.zone01dakar.sn/forum-rest-api/lib/errors"

	"learn.zone01dakar.sn/forum-rest-api/models"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

type Session struct{}

func SessionService() Session {
	return Session{}
}

func (s *Session) Create(response *lib.Response, db lib.DB, userID string) (models.Session, error) {
	var sqlService = service.SqlService(db.Instance)
	var queryInstance = core.NewQuery()
	var session = models.Session{
		Token:          lib.TokenGenerator().String(),
		ExpirationDate: time.Now().Add(24 * time.Hour),
		UserID:         userID,
		CreationDate:   time.Now(),
	}

	insertQuery, _ := queryInstance.INSERT("token", "expires", "user_id").FROMTABLE("Session").Build()
	updateQuery := "UPDATE Session SET token = ?, expires = ? WHERE user_id = ?"

	sqlErr := sqlService.Create(insertQuery, lib.Slicer(session, false)...)

	_, statuscode := errHandler.SqlError(sqlErr, []string{})

	// Check if the user already got a session if not it creates a new one
	if statuscode == 409 {
		if updateERR := sqlService.Update(updateQuery, session.Token, session.ExpirationDate, session.UserID); updateERR != nil {
			return models.Session{}, updateERR
		}
	}
	return session, nil

}
