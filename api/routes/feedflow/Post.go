package feedflow

import (
	"net/http"
	"strings"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/lib"
	"learn.zone01dakar.sn/forum-rest-api/models"
	service "learn.zone01dakar.sn/forum-rest-api/service/CRUD"
)

type Post struct{}

func (p Post) Route(app *core.App) {
	app.POST("/posts", p.Create)
	// app.GET("/posts/", p.GetPostByID)
	app.GET("/posts", p.GetAllPost)

}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.Post
	post := CreateFeed{
		Credentials:   &credentials,
		Fields:        []string{"image", "title", "content", "author_id"},
		Table:         "Post",
		ItemsToInsert: []interface{}{&credentials.Image, &credentials.Title, &credentials.Content, &credentials.AuthorID},
		SumittedData: map[string]interface{}{
			"title":     &credentials.Title,
			"content":   &credentials.Content,
			"author_id": &credentials.AuthorID,
		},
	}

	post.Create(w, r, &response)
	lib.ResponseFormatter(w, response)
}

func (p Post) GetAllPost(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	db, _ := r.Context().Value("db").(lib.DB)
	posts := []models.Post{}
	postsID := []string{}

	var sqlService = service.SqlService(db.Instance)
	var queryPosts, _ = core.NewQuery().SELECT("id").FROMTABLE("Post").Build()
	sqlService.SelectAllForPrimitive(queryPosts, "", &postsID)

	for _, id := range postsID {
		post := p.P(w, r, &response, id)
		posts = append(posts, post)
	}

	response.Data = posts
	lib.ResponseFormatter(w, response)
}

func (p Post) GetPostByID(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok", Data: "The main character"}
	id := strings.TrimPrefix(r.URL.Path, r.URL.Path)
	p.P(w, r, &response, id)
	lib.ResponseFormatter(w, response)
}

func (p Post) P(w http.ResponseWriter, r *http.Request, response *lib.Response, id string) models.Post {
	post := models.Post{}

	feed := GetFeed{
		Table:       "Post",
		Credentials: &post,
		Dest:        []interface{}{&post.ID, &post.Image, &post.Title, &post.Content, &post.AuthorID, &post.CreationDate},
		Items: []Items{
			{Item: "id", EntrieID: "post_id", Table: "Comments", Dest: &post.Comments},
			{Item: "id", EntrieID: "entries_id", Table: "Postlikes", Dest: &post.Likes},
			{Item: "id", EntrieID: "entries_id", Table: "Postdislikes", Dest: &post.DisLikes},
		},
		UserInfo: []interface{}{&post.FirstName, &post.LastName, &post.Username, &post.Avatar},
		AuthorID: &post.AuthorID,
	}

	// var queryCategory, _ = core.NewQuery().SELECT("category_id").WHERE("post_id = ?").FROMTABLE("PostCategories").Build()
	feed.GetByID(w, r, response, id)
	return post
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Implement logic to update a post by ID
	response := lib.Response{Code: 200, Message: "ok"}
	// db, _ := r.Context().Value("db").(lib.DB)

	lib.ResponseFormatter(w, response)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Implement logic to delete a post by ID
	response := lib.Response{Code: 200, Message: "ok"}
	// db, _ := r.Context().Value("db").(lib.DB)

	lib.ResponseFormatter(w, response)
}

// {
//     "FirstName": "Test",
//     "LastName": "Test",
//     "Email": "abdouaziznjay@gmail.com",
//     "Username": "aziz",
//     "Bio": "Software engineer at Zone01 Dakar",
//     "Avatar": "http://source.unsplash.com/50x50",
//     "Password": "1234"
// }

// {
// 	"Image":   "any",
// 	"Title":   "Testing post feature",
// 	"Content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Temporibus, illum porro explicabo earum at quod distinctio, laudantium velit voluptatibus provident quidem ea dolore delectus ipsam tenetur. Rem blanditiis dolores, culpa praesentium facilis, tenetur nihil porro numquam ipsam deserunt sapiente odio!",
// 	"AuthorID":  "1",

// 	"Likes":    [],
// 	"DisLikes": [],

// 	"Comments": [],
// 	"Category": [],

// 	"Username":  "aziz",
// 	"FirstName": "aziz",
// 	"LastName":  "ndiaye",
// 	"Avatar":    "string"
// }

// func (p Post) GetAllPost(w http.ResponseWriter, r *http.Request) {
// 	response := lib.Response{Code: 200, Message: "ok"}
// 	db, _ := r.Context().Value("db").(lib.DB)
// 	posts := []models.Post{}

// 	var sqlService = service.SqlService(db.Instance)

// 	var queryPosts, _ = core.NewQuery().SELECT("*").FROMTABLE("Post").Build()
// 	err := sqlService.SelectAllForStruct(queryPosts, "", &posts, []string{"ID", "Image", "Title", "Content", "AuthorID", "CreationDate"})
// 	if err != nil {
// 		message, statusCode := errors.SqlError(err)
// 		errors.ErrorWriter(&response, message, statusCode)
// 	}

// 	errCh := make(chan error, 4)
// 	authorID := make(chan string, len(posts))
// 	wg := sync.WaitGroup{}
// 	wg.Add(4 * len(posts))

// 	for _, post := range posts {

// 		go func(post models.Post) {
// 			defer fmt.Println(post.Comments)

// 			feed := GetFeed{
// 				Items: []Items{
// 					{Item: "id", EntrieID: "post_id", Table: "Comments", Dest: &post.Comments},
// 					{Item: "id", EntrieID: "entries_id", Table: "Postlikes", Dest: &post.Likes},
// 					{Item: "id", EntrieID: "entries_id", Table: "Postdislikes", Dest: &post.DisLikes},
// 				},
// 				UserInfo: []interface{}{&post.FirstName, &post.LastName, &post.Username, &post.Avatar},
// 			}

// 			authorID <- post.AuthorID
// 			feed.GetOtherItemsPost(post.ID, &response, sqlService, errCh, authorID, &wg)

// 		}(post)

// 	}

// 	go func() {
// 		wg.Wait()
// 		close(errCh)
// 		close(authorID)
// 	}()

// 	for err := range errCh {
// 		if err != nil {
// 			lib.ResponseFormatter(w, response)
// 			return
// 		}
// 	}

// 	response.Data = posts
// 	lib.ResponseFormatter(w, response)
// }
