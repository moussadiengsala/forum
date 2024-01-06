package feedflow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	errors "learn.zone01dakar.sn/forum-rest-api/lib/errors"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

type CreateFeed struct {
	Fields        []string
	Table         string
	Credentials   interface{}
	ItemsToInsert []interface{}
}

func (f *CreateFeed) Create(w http.ResponseWriter, r *http.Request, response *lib.Response) {

	if err := json.NewDecoder(r.Body).Decode(&f.Credentials); err != nil {
		errors.ErrorWriter(response, "Error getting informations to perform this operation!!!", http.StatusBadRequest)
		return
	}

	validators := core.Validators{}
	if errValidator := validators.ValidatorService(f.Credentials); errValidator != nil {
		errors.ErrorWriter(response, errValidator.Error(), http.StatusBadRequest)
		return
	}

	db, _ := r.Context().Value("db").(lib.DB)

	var sqlService = service.SqlService(db.Instance)
	var query, _ = core.NewQuery().INSERT(f.Fields...).FROMTABLE(f.Table).Build()

	insertPostErr := sqlService.Create(query, f.ItemsToInsert...)

	if insertPostErr != nil {
		message, statusCode := errors.SqlError(insertPostErr, []string{"post", "comment"})
		errors.ErrorWriter(response, message, statusCode)
		return
	}

	response.Data = f.Credentials
}

type Items struct {
	Item     string
	EntrieID string
	Table    string
	Dest     interface{}
}

type GetFeed struct {
	Table       string
	Credentials interface{}
	Dest        []interface{}
	Items       []Items
	UserInfo    []interface{}
	AuthorID    *string
}

func (f GetFeed) GetByID(w http.ResponseWriter, r *http.Request, response *lib.Response, id string) {
	db, _ := r.Context().Value("db").(lib.DB)

	var wg sync.WaitGroup
	authorID := make(chan string, 1)
	errCh := make(chan error, len(f.Items)+1)
	wg.Add(len(f.Items) + 2)

	var sqlService = service.SqlService(db.Instance)

	go func() {
		defer wg.Done()
		var queryPosts, _ = core.NewQuery().SELECT("*").WHERE("id = ?").FROMTABLE(f.Table).Build()
		err := sqlService.SelectSingle(queryPosts, []interface{}{id}, f.Dest...)
		if err != nil {
			message, statusCode := errors.SqlError(err, []string{"post", "comment"})
			errors.ErrorWriter(response, message, statusCode)
			errCh <- err
		}
		authorID <- *f.AuthorID
	}()

	f.GetOtherItemsPost(id, response, sqlService, errCh, authorID, &wg)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return
		}
	}

	response.Data = f.Credentials
}

func (f *GetFeed) GetOtherItemsPost(id string, response *lib.Response, sqlService service.DBService, errCh chan error, authorID chan string, wg *sync.WaitGroup) {
	for _, item := range f.Items {
		go func(item Items) {
			defer wg.Done()
			var query, _ = core.NewQuery().SELECT(item.Item).WHERE(fmt.Sprintf("%s = ?", item.EntrieID)).FROMTABLE(item.Table).Build()
			err := sqlService.SelectAllForPrimitive(query, id, item.Dest)

			if err != nil {
				message, statusCode := errors.SqlError(err, []string{"post", "comment"})
				errors.ErrorWriter(response, message, statusCode)
				errCh <- err
			}
		}(item)
	}

	go func() {
		defer wg.Done()
		fields := []string{"first_name", "last_name", "username", "avatar"}
		var queryUser, _ = core.NewQuery().SELECT(fields...).WHERE("id = ?").FROMTABLE("User").Build()
		err := sqlService.SelectSingle(queryUser, []interface{}{<-authorID}, f.UserInfo...)
		if err != nil {
			message, statusCode := errors.SqlError(err, []string{"post", "comment"})
			errors.ErrorWriter(response, message, statusCode)
			errCh <- err

		}
	}()

}
