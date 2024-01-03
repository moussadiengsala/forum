package service

import (
	"context"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errors "learn.zone01dakar.sn/forum-rest-api/lib/errors"
)

func DBMiddleware(next lib.Handler) lib.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		response := lib.Response{}
		DB := lib.DB{Instance: nil, Err: nil}

		database := internals.Config{
			Driver: "sqlite3",
			Name:   "forum.db",
		}
		/* Check if a database connection is already opened
		to avoid opening multiple databases connections */
		if DB.Instance == nil {
			DB.Instance, DB.Err = database.Init()
			// Throw the error page when we have a database issue
			if DB.Err != nil {
				errors.ErrorWriter(&response,DB.Err.Error(),http.StatusInternalServerError)
				lib.ResponseFormatter(w, response)
				return
			}
			internals.TablesCreation(DB.Instance)
		}

		var ctx = context.Background()
		ctx = context.WithValue(r.Context(), "db", DB)
		next(w, r.WithContext(ctx))
	}
}
