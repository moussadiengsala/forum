package externals

import "encoding/json"

var GoogleAuth = OAuthConfig{
	Name:         "GOOGLE",
	ClientID:     "341396623977-gn03cld9td2djqmia0allkb9jv37uvm5.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURI:  "http://localhost:8080/auth/register",
	Scope:        []string{"email", "profile", "openid"},
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	UserInfoURL:  "https://www.googleapis.com/oauth2/v1/userinfo",
}

func ParseAccessTokenGoogle(body []byte) (string, error) {
	var token map[string]interface{}
	err := json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}
	return token["access_token"].(string), nil
}
