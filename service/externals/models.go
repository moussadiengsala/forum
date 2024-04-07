package externals

import (
	"fmt"
	"os"
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

func (o *OAuthConfig) SetAuthURL(endpoint string) string {
	o.ClientID = os.Getenv(fmt.Sprintf("%s_CLIENT_ID", o.Name))
	o.ClientSecret = os.Getenv(fmt.Sprintf("%s_CLIENT_SECRET", o.Name))
	o.RedirectURI = fmt.Sprintf("http://localhost:8080/auth/%s", endpoint)
	params := BuildURL(o.AuthURLQuery())
	// fmt.Print(params)
	return fmt.Sprintf("%s?%s", o.AuthURL, params.Encode())
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
