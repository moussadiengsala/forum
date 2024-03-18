package handlers

import (
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"golang-rest-api-starter/service/externals"
	"log"
	"net/http"
)

func HandleAuthentificationProvider(r *http.Request, response *model.Reponse, db *model.DB) (bool, model.User) {
	user := model.User{}
	var codeAuthProvider = r.URL.Query().Get("code")
	var provider, ok = externals.GetProvider(r)

	if codeAuthProvider != "" && ok && r.Method == "GET" {
		var authProviderData, err = externals.GetUserInfos(codeAuthProvider, provider)
		if err != nil {
			helpers.ErrorWriter(response, "Error getting user info: "+err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return false, user
		}

		if provider.Name == "GOOGLE" {
			user = externals.GetUserInfosNeededGoogle(authProviderData)
		} else if provider.Name == "GITHUB" {
			user = externals.GetUserInfosNeededGithub(authProviderData)
		}

		if err != nil {
			helpers.ErrorWriter(response, "Error getting user info: "+err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return false, user
		}

		return true, user
	}
	return false, user
}
