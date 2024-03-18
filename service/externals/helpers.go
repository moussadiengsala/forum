package externals

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func BuildURL(URLQuery map[string]string) url.Values {
	var params = url.Values{}
	for k, v := range URLQuery {
		params.Add(k, v)
	}

	return params
}

func GetProvider(r *http.Request) (OAuthConfig, bool) {
	var providers = map[string]OAuthConfig{
		"google": GoogleAuth,
		"github": GithubAuth,
	}

	var provider, ok = providers[r.URL.Query().Get("state")]
	return provider, ok
}

func GetAccessToken(authCode string, provider OAuthConfig) (string, error) {
	// provider.SetAuthCode(authCode)
	provider.AuthCode = authCode
	var params = BuildURL(provider.TokenURLQuery())
	// Exchaning authCode to access Token.
	respToAccessToken, errToAccessToken := http.PostForm(provider.TokenURL, params)
	// fmt.Println(respToAccessToken)
	if errToAccessToken != nil {
		return "", errors.New(errToAccessToken.Error())
	}
	defer respToAccessToken.Body.Close()

	// Parsing the response token
	body, errResponse := io.ReadAll(respToAccessToken.Body)
	if errResponse != nil {
		return "", errors.New(errResponse.Error())
	}

	var (
		token      string
		errParsing error
	)

	if provider.Name == "GOOGLE" {
		token, errParsing = ParseAccessTokenGoogle(body)
	} else if provider.Name == "GITHUB" {
		token, errParsing = ParseAccessTokenGithub(body)
	}

	if errParsing != nil {
		return "", errors.New(errParsing.Error())
	}

	return token, nil
}

func GetUserInfos(authCode string, provider OAuthConfig) (map[string]interface{}, error) {
	var userInfo map[string]interface{}
	var token, err = GetAccessToken(authCode, provider)

	// Request to access user's info by providing the access token.
	requestUserInfo, errGettingUserInfo := http.NewRequest("GET", provider.UserInfoURL, nil)
	if errGettingUserInfo != nil {
		return nil, errors.New(errGettingUserInfo.Error())
	}
	requestUserInfo.Header.Set("Authorization", "Bearer "+token)

	var client = &http.Client{}
	response, err := client.Do(requestUserInfo)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer response.Body.Close()

	dataUser, errResponse := io.ReadAll(response.Body)
	if errResponse != nil {
		return nil, errors.New(errResponse.Error())
	}

	errUnmarshal := json.Unmarshal(dataUser, &userInfo)
	if errUnmarshal != nil {
		return nil, errors.New(errUnmarshal.Error())
	}

	return userInfo, nil
}
