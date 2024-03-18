package handlers

import (
	"errors"
	internals "golang-rest-api-starter/internals/config/database"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/externals"
	"net/http"
)

type Auth struct{}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	db, errInstance := r.Context().Value("db").(model.DB)
	session := Session{}
	internals.TablesCreation(db.Instance)

	// if got error getting database instance.
	if !errInstance {
		helpers.ErrorThrower(errors.New("Error getting database instance"), "Something went wrong", http.StatusInternalServerError, w, r)
		return
	}

	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}

	if r.Method == http.MethodPost {
		datas := helpers.GetFormData(r)
		//server side validation 🔒
		helpers.ValidateForm(datas, &response)
		var err = helpers.IsRequiredFeildsExits(datas, "email", "password")
		if err != nil {
			helpers.ErrorWriter(&response, err.Error(), http.StatusBadRequest)
		}
		//
		if response.StatusCode == http.StatusOK {
			var usr = helpers.VerifyUserInfo(&response, &db, true, "email", datas)
			session.CreateSession(w, r, &response, &db, usr)
			return
		}
	}

	if isAuthenticated, user := HandleAuthentificationProvider(r, &response, &db); isAuthenticated {

		var datas = map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		}
		var usr = helpers.VerifyUserInfo(&response, &db, false, "email", datas)
		session.CreateSession(w, r, &response, &db, usr)
		// return
	}

	response.Data = map[string]string{
		"GOOGLE_AUTH": externals.GoogleAuth.SetAuthURL("login"),
		"GITHUB_AUTH": externals.GithubAuth.SetAuthURL("login"),
	}
	helpers.ResponseFormatter(response, "login", w, r, response.StatusCode)
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	db, _ := r.Context().Value("db").(model.DB)
	// Get the submitted data from the client
	datas := helpers.GetFormData(r)
	// Default response sent by the server
	response := model.Reponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
		HasError:   true,
	}

	if r.Method == http.MethodPost {
		//server side validation 🔒
		helpers.ValidateForm(datas, &response)
		var err = helpers.IsRequiredFeildsExits(datas, "fname", "lname", "email", "username", "password")
		if err != nil {
			helpers.ErrorWriter(&response, err.Error(), http.StatusBadRequest)
		}
		bio, ok := datas["bio"]
		if ok {
			if len(bio.(string)) > 255 {
				helpers.ErrorWriter(&response, "Your bio cannot exceed 255 characters.", http.StatusBadRequest)
			}
		}
		// if datas are valid ✅
		if response.StatusCode == http.StatusOK {
			randomAvatar := helpers.AvatarGenerator()
			hashedPassword := helpers.PasswordEncrypter(datas["password"].(string))

			helpers.SaveUserInfo(&db, &response, datas["fname"], datas["lname"], datas["email"], datas["username"], bio, randomAvatar, hashedPassword)
		}
	}

	var isRequestAuthenticatedProvider, user = HandleAuthentificationProvider(r, &response, &db)

	if isRequestAuthenticatedProvider {
		helpers.SaveUserInfo(&db, &response, user.FirstName, user.LastName, user.Email, user.Username, user.Bio, user.Avatar, "")
	}

	response.Data = map[string]string{
		"GOOGLE_AUTH": externals.GoogleAuth.SetAuthURL("register"),
		"GITHUB_AUTH": externals.GithubAuth.SetAuthURL("register"),
	}

	// Send response to the client
	helpers.ResponseFormatter(response, "register", w, r, response.StatusCode)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	db, _ := r.Context().Value("db").(model.DB)
	userID, _ := r.Context().Value("userID").(string)
	db.Instance.Exec("DELETE FROM Sessions WHERE token = ?", userID)
}
