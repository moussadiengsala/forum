package externals

import (
	"fmt"
	model "golang-rest-api-starter/models"
	"net/url"
)

var GithubAuth = OAuthConfig{
	Name:         "GITHUB",
	ClientID:     "",
	ClientSecret: "",
	Scope:        []string{"user"},
	RedirectURI:  "http://localhost:8080/auth/",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
	UserInfoURL:  "https://api.github.com/user",
	State:        "github",
}

func ParseAccessTokenGithub(body []byte) (string, error) {
	var value, err = url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}
	return value.Get("access_token"), nil
}

func GetUserInfosNeededGithub(data map[string]interface{}) model.User {
	var user model.User

	var converdata = func(data interface{}) string {
		if data == nil {
			return ""
		}
		var value, ok = data.(string)
		if !ok {
			return ""
		}
		return value
	}

	user.FirstName = converdata(data["name"])
	user.Username = fmt.Sprintf("%s.%f", converdata(data["login"]), data["id"])
	user.Avatar = converdata(data["avatar_url"])
	user.Bio = converdata(data["bio"])
	user.Email = fmt.Sprintf("%s%s@%s", converdata(data["login"]), data["id"], "github.com")

	return user
}
