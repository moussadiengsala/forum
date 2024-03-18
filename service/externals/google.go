package externals

import (
	"encoding/json"
	"errors"
	"fmt"
	model "golang-rest-api-starter/models"
	"os"
	"strings"
)

var GoogleAuth = OAuthConfig{
	Name:         "GOOGLE",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURI:  "http://localhost:8080/auth/login",
	Scope:        []string{"email", "profile", "openid"},
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	UserInfoURL:  "https://www.googleapis.com/oauth2/v1/userinfo",
	State:        "google",
}

func ParseAccessTokenGoogle(body []byte) (string, error) {
	var token map[string]interface{}
	if err := json.Unmarshal(body, &token); err != nil {
		return "", err
	}

	// Check if the access token field is empty
	val, ok := token["access_token"].(string)
	if !ok {
		return "", errors.New("access token not found in JSON response")
	}

	return val, nil
}

func GetUserInfosNeededGoogle(data map[string]interface{}) model.User {
	user := model.User{}
	id := ""
	fields := []struct {
		key   string
		field *string
	}{
		{"given_name", &user.FirstName},
		{"family_name", &user.LastName},
		{"picture", &user.Avatar},
		{"email", &user.Email},
		{"id", &id},
	}

	for _, f := range fields {
		val, ok := data[f.key].(string)
		if !ok {
			val = ""
		}
		*f.field = val
	}

	user.Username = fmt.Sprintf("%s.%s", strings.Split(user.Email, "@")[0], id)

	return user
}

// func GetUserInfosNeededGoogle(data map[string]interface{}) model.User {
// 	var user model.User

// 	user.FirstName = data["given_name"].(string)
// 	user.LastName = data["family_name"].(string)
// 	user.Avatar = data["picture"].(string)
// 	user.Email = data["email"].(string)
// 	user.Username = fmt.Sprintf("%s.%s", strings.Split(user.Email, "@")[0], data["id"])

// 	return user
// }

// Data:  map[
// 	email:moizadieng@gmail.com
// 	family_name:Dieng
// 	given_name:Moussa
// 	id:100962570289848056856
// 	locale:fr
// 	name:Moussa Dieng
// 	picture:https://lh3.googleusercontent.com/a/ACg8ocLgDbvkbCJ_1FJeKqVuszaLU237xoBLFnPnsTKAHXds58k=s96-c
// 	verified_email:true]
