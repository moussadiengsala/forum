package feedflow

import (
	"net/http"
	"strings"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	"learn.zone01dakar.sn/forum-rest-api/models"
)

type Comment struct{}

func (p Comment) Route(app *core.App) {
	app.POST("/comments", p.Create)
	app.GET("/comments/", p.GetCommentByID)
}

func (c Comment) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	var credentials models.Comments
	post := CreateFeed{
		Credentials:   &credentials,
		Fields:        []string{"content", "author_id", "post_id"},
		Table:         "Comments",
		ItemsToInsert: []interface{}{&credentials.Content, &credentials.AuthorID, &credentials.PostID},
	}
	post.Create(w, r, &response)
	lib.ResponseFormatter(w, response)
}

func (c Comment) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	comment := models.Comments{}
	id := strings.TrimPrefix(r.URL.Path, r.URL.Path)

	feed := GetFeed{
		Table:       "Comments",
		Credentials: &comment,
		Dest:        []interface{}{&comment.ID, &comment.Content, &comment.AuthorID, &comment.PostID, &comment.CreationDate},
		Items: []Items{
			{Item: "id", EntrieID: "entries_id", Table: "Commentlikes", Dest: &comment.Likes},
			{Item: "id", EntrieID: "entries_id", Table: "Commentdislikes", Dest: &comment.DisLikes},
		},
		UserInfo: []interface{}{&comment.User.FirstName, &comment.User.LastName, &comment.User.Username, &comment.User.Avatar},
		AuthorID: &comment.AuthorID,
	}

	feed.GetByID(w, r, &response, id)
	lib.ResponseFormatter(w, response)
}

// {
// 	"Content":  "Fuck you",
// 	"AuthorID": "1",
// 	"PostID":   "1",
// }
