package middleware

import (
	"fmt"
	databaseConfig "golang-rest-api-starter/internals/config/database"
	model "golang-rest-api-starter/models"
	"net/http"
)

func DatabaseMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var DB = model.DB{Instance: nil, Err: nil}

		var database = databaseConfig.Config{
			Driver: "sqlite3",
			Name:   "forum.db",
		}

		DB.Instance, DB.Err = database.Init()
		if DB.Err != nil {
			fmt.Println(DB.Err)
		}
		databaseConfig.TablesCreation(DB.Instance)

		handler(w, r)
	}
}
