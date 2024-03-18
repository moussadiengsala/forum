package helpers

import (
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/crud"
	"log"
	"net/http"

	"github.com/mattn/go-sqlite3"
)

// Save the user's info in the database while registering
func SaveUserInfo(db *model.DB, response *model.Reponse, value ...interface{}) {
	columnsToInsert := []string{"first_name", "last_name", "email", "username", "bio", "avatar", "password"}
	_, insertionErr := crud.Insert(db.Instance, "user", columnsToInsert, value...)

	if insertionErr != nil {
		log.Println(insertionErr)
		if sqliteErr, ok := insertionErr.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			ErrorWriter(response, "*This email gituhh or username already exists. Please use a different one", http.StatusConflict)
			log.Println(sqliteErr)
		} else {
			ErrorWriter(response, "Something went wrong. If the problem persists, please contact us", http.StatusInternalServerError)
			log.Println(insertionErr)
		}
	}

	if response.StatusCode == http.StatusOK {
		response.StatusCode = 201
		response.HasError = false
	}
}
