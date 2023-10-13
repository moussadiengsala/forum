package externals

import (
	"fmt"
	"golang-rest-api-starter/internals/helpers"
	"strings"
)

type OAuthConfig struct {
	Name         string
	ClientID     string
	ClientSecret string
	Scope        []string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	UserInfoURL  string
	State        string
	AuthCode     string
}

func (o *OAuthConfig) SetAuthURL() string {
	o.State = helpers.IDGenerator().String()

	params := BuildURL(o.AuthURLQuery())
	o.AuthURL = fmt.Sprintf("%s?%s", o.AuthURL, params.Encode())
	return o.AuthURL
}

func (o *OAuthConfig) SetAuthCode(authCode string) {
	o.AuthCode = authCode
}

func (o OAuthConfig) AuthURLQuery() map[string]string {
	return map[string]string{
		"client_id":     o.ClientID,
		"redirect_uri":  o.RedirectURI,
		"scope":         strings.Join(o.Scope, " "),
		"response_type": "code",
		"state":         o.State,
	}
}

func (o OAuthConfig) TokenURLQuery() map[string]string {
	return map[string]string{
		"client_id":     o.ClientID,
		"client_secret": o.ClientSecret,
		"code":          o.AuthCode,
		"redirect_uri":  o.RedirectURI,
		"grant_type":    "authorization_code",
	}
}
