package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/internals/config/session"
	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errors "learn.zone01dakar.sn/forum-rest-api/lib/errors"
	"learn.zone01dakar.sn/forum-rest-api/models"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

type Auth struct{}

func (a Auth) Route(app *core.App) {
	app.POST("/auth/signin", a.SignIn)
	app.POST("/auth/signup", a.Register)
}

func (a Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials lib.Credentials

	// if credentials.Identifiers == "" || credentials.Password == "" that means he doesn't fulfill one of them.
	errGettingCredential := json.NewDecoder(r.Body).Decode(&credentials)
	if errGettingCredential != nil || credentials.Identifiers == "" || credentials.Password == "" {
		errors.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	// By default we suppose the user login with his username
	var identifiers = "username"

	validators := core.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		errors.ErrorWriter(&response, errValidator.Error(), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value("db").(lib.DB)
	var session = session.SessionService()
	var sqlService = service.SqlService(db.Instance)
	var query, _ = core.NewQuery().
		SELECT("id", "first_name", "last_name", "email", "username", "bio", "avatar", "password").
		FROMTABLE("User").WHERE(fmt.Sprintf("%s = ? ", identifiers)).
		Build()

	var userData models.User

	err := sqlService.SelectSingle(query, []interface{}{credentials.Identifiers}, &userData.ID, &userData.FirstName, &userData.LastName, &userData.Email, &userData.Username, &userData.Bio, &userData.Avatar, &userData.Password)
	if err != nil {
		message, statuscode := errors.SqlError(err, []string{"email", "username"})
		errors.ErrorWriter(&response, message, statuscode)
		lib.ResponseFormatter(w, response)
		return
	}

	if !lib.PasswordDecrypter(userData.Password, credentials.Password) {
		errors.ErrorWriter(&response, "*Incorrect password!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	response.Data = userData
	sessions, err := session.Create(&response, db, userData.ID)
	if err != nil {
		errors.ErrorWriter(&response, err.Error(), http.StatusInternalServerError)
		lib.ResponseFormatter(w, response)
		return
	}
	userData.Password = ""
	payload := lib.Payload{
		User:    userData, 
		Session: sessions,
	}
	response.Data = payload

	lib.ResponseFormatter(w, response)
}

func (a Auth) Register(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var userData models.User

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		errors.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields !", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	validators := core.Validators{}
	if errValidator := validators.ValidatorService(userData); errValidator != nil {
		errors.ErrorWriter(&response, errValidator.Error(), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value("db").(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	fields := []string{"first_name", "last_name", "email", "username", "bio", "avatar", "password"}
	var query, _ = core.NewQuery().INSERT(fields...).FROMTABLE("User").Build()

	userData.Password = lib.PasswordEncrypter(userData.Password)
	if err := sqlService.Create(query, lib.Slicer(userData, false)...); err != nil {
		message, code := errors.SqlError(err, []string{"email", "username"})
		errors.ErrorWriter(&response, message, code)
	}

	lib.ResponseFormatter(w, response)
}
