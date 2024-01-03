package service

import (
	"net/http"
	"strings"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errHandler "learn.zone01dakar.sn/forum-rest-api/lib/errors"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

func AuthMiddleware(next lib.Handler) lib.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/auth/") {
			next(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		db, _ := r.Context().Value("db").(lib.DB)
		var query, _ = core.NewQuery().
			SELECT("token").
			FROMTABLE("Session").WHERE("token = ? ").
			Build()

		var token string
		var sqlService = service.SqlService(db.Instance)
		if err := sqlService.SelectSingle(query, []interface{}{authorizationHeader}, &token); err != nil {
			_, statuscode := errHandler.SqlError(err, []string{"token"})
			if statuscode != http.StatusOK {
				lib.ResponseFormatter(w, lib.Response{Code: http.StatusForbidden,
					Message: "You are not allowed to access to this ressource.Please login and try again."})
				return
			}
		}

		next(w, r)

	}
}
