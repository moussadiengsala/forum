package feedflow

import (
	"encoding/json"
	"fmt"
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errors "learn.zone01dakar.sn/forum-rest-api/lib/errors"
	"learn.zone01dakar.sn/forum-rest-api/models"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

type Reaction struct {
	Action map[string][]string
}

func (re *Reaction) SetAction() {
	re.Action = make(map[string][]string)
	re.Action["postlikes"] = []string{"Postlikes", "Postdislikes"}
	re.Action["postdislikes"] = []string{"Postdislikes", "Postlikes"}
	re.Action["commentlikes"] = []string{"Commentlikes", "Commentdislikes"}
	re.Action["commentdislikes"] = []string{"Commentdislikes", "Commentlikes"}
}

func (r *Reaction) Route(app *core.App) {
	r.SetAction()
	app.POST("/reactions", r.Create)
}

func (re Reaction) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	var credentials models.Reaction
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		errors.ErrorWriter(&response, "Error getting informations to perform this operation!!!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	// cre := map[string]interface{}{
	// 	"author_id":  credentials.AuthorID,
	// 	"entries_id": credentials.EntriesID,
	// 	"action":     credentials.Action,
	// }

	// if errValidator := validators.ValidatorService(cre); errValidator != nil {
	// 	errors.ErrorWriter(&response, errValidator.Error(), http.StatusBadRequest)
	// 	lib.ResponseFormatter(w, response)
	// 	return
	// }

	tables, _ := re.Action[credentials.Action]
	db, _ := r.Context().Value("db").(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	err := re.CheckAndDeleteReaction(sqlService, tables[0], []interface{}{credentials.AuthorID, credentials.EntriesID}, &credentials.ID)
	_, statusCode := errors.SqlError(err, []string{})

	if statusCode == 404 {
		q, _ := core.NewQuery().INSERT("author_id", "entries_id").FROMTABLE(tables[0]).Build()
		if errCreate := sqlService.Create(q, credentials.AuthorID, credentials.EntriesID); errCreate != nil {
			message, statusCode := errors.SqlError(errCreate, []string{})
			errors.ErrorWriter(&response, message, statusCode)
			lib.ResponseFormatter(w, response)
			return
		}

		errI := re.CheckAndDeleteReaction(sqlService, tables[1], []interface{}{credentials.AuthorID, credentials.EntriesID}, &credentials.ID)
		message, statusCode := errors.SqlError(errI, []string{})
		if errI != nil && statusCode != 404 {
			errors.ErrorWriter(&response, message, statusCode)
			lib.ResponseFormatter(w, response)
			return
		}
	}

	response.Data = credentials
	lib.ResponseFormatter(w, response)
}

func (re Reaction) CheckAndDeleteReaction(sqlService service.DBService, table string, values []interface{}, dest *string) error {
	checkQuery, _ := core.NewQuery().SELECT("id").WHERE("author_id = ? AND entries_id = ?").FROMTABLE(table).Build()
	deletequery := fmt.Sprintf("DELETE FROM %s WHERE author_id = ? AND entries_id = ?;", table)

	err := sqlService.SelectSingle(checkQuery, values, dest)
	if err == nil {
		err = sqlService.Delete(deletequery, values...)
	}
	return err
}

func (c Reaction) Get(ws http.ResponseWriter, r *http.Request) {

}

// {
//     "AuthorID":  "1",
//     "EntriesID": "3",
// 	"Action": "postlikes"
// }
