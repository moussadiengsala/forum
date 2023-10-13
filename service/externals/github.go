package externals

import "net/url"

var GithubAuth = OAuthConfig{
	Name:         "GITHUB",
	ClientID:     "3a2bb5d588e9c9b5951b",
	ClientSecret: "b4cce0c5de9a046d43c792ca5ce84ef0b1dfa08b",
	Scope:        []string{"user"},
	RedirectURI:  "http://localhost:8080/auth/register",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
	UserInfoURL:  "https://api.github.com/user",
}

func ParseAccessTokenGithub(body []byte) (string, error) {
	var value, err = url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}
	return value.Get("access_token"), nil
}

// func GetUserInfoG(r *http.Request, authCode, state string) (ProvideRsponse, error) {
// 	var responseProvider ProvideRsponse
// 	GithubAuth.SetAuthCode(authCode)
// 	var params = BuildURL(GithubTokenURLQuery())
// 	// fmt.Println("authCode", authCode)
// 	// fmt.Println(params)
// 	// Exchaning authCode to access Token.
// 	respToAccessToken, errToAccessToken := http.Post(GithubAuth.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
// 	// respToAccessToken, errToAccessToken := http.PostForm(GithubAuth.TokenURL, params)
// 	if errToAccessToken != nil {
// 		return responseProvider, errors.New(errToAccessToken.Error())
// 	}
// 	defer respToAccessToken.Body.Close()

// 	// Parsing the response token
// 	// fmt.Println("Response Token", respToAccessToken)
// 	body, _ := io.ReadAll(respToAccessToken.Body)
// 	value, _ := url.ParseQuery(string(body))
// 	fmt.Println(value)
// 	fmt.Println("Access Token:", value.Get("access_token"))
// 	// fmt.Println("acessToken", value.Get("access_token"))
// 	token := value.Get("access_token")

// 	// Request to access user's info by providing the access token.
// 	requestUserInfo, errGettingUserInfo := http.NewRequest("GET", GithubAuth.UserInfoURL, nil)
// 	if errGettingUserInfo != nil {
// 		return responseProvider, errors.New(errGettingUserInfo.Error())
// 	}
// 	requestUserInfo.Header.Set("Authorization", "Bearer "+token)

// 	var client = &http.Client{}
// 	response, err := client.Do(requestUserInfo)
// 	if err != nil {
// 		return responseProvider, errors.New(err.Error())
// 	}
// 	defer response.Body.Close()
// 	// fmt.Println(response)

// 	var userInfo map[string]interface{}
// 	b, _ := io.ReadAll(response.Body)
// 	_ = json.Unmarshal(b, &userInfo)

// 	// q, _ := url.ParseQuery(string(b))
// 	// fmt.Println(q)
// 	fmt.Println("id: ", userInfo["id"])
// 	fmt.Println("username: ", userInfo["login"])
// 	fmt.Println("avatar: ", userInfo["avatar_url"])
// 	fmt.Println("email: ", userInfo["email"])
// 	fmt.Println("bio: ", userInfo["bio"])

// 	var errParsingUserInfo = json.NewDecoder(response.Body).Decode(&responseProvider)
// 	if errParsingUserInfo != nil {
// 		return responseProvider, errors.New(errParsingUserInfo.Error())
// 	}

// 	return responseProvider, nil
// }

// map[{
// "login":"moussadiengsala",
// "id":63497230,
// "node_id":"MDQ6VXNlcjYzNDk3MjMw",
// "avatar_url":"https://avatars.githubusercontent.com/u/63497230?v:[4",
// "gravatar_id":"","url":"https://api.github.com/users/moussadiengsala",
// "html_url":"https://github.com/moussadiengsala",
// "followers_url":"https://api.github.com/users/moussadiengsala/followers",
// "following_url":"https://api.github.com/users/moussadiengsala/following{/other_user}",
// "gists_url":"https://api.github.com/users/moussadiengsala/gists{/gist_id}",
// "starred_url":"https://api.github.com/users/moussadiengsala/starred{/owner}{/repo}",
// "subscriptions_url":"https://api.github.com/users/moussadiengsala/subscriptions",
// "organizations_url":"https://api.github.com/users/moussadiengsala/orgs",
// "repos_url":"https://api.github.com/users/moussadiengsala/repos",
// "events_url":"https://api.github.com/users/moussadiengsala/events{/privacy}",
// "received_events_url":"https://api.github.com/users/moussadiengsala/received_events",
// "type":"User",
// "site_admin":false,
// "name":null,
// "company":null,
// "blog":"",
// "location":null,
// "email":null,
// "hireable":null,
// "bio":null,
// "twitter_username":null,
// "public_repos":5,
// "public_gists":0,
// "followers":0,
// "following":0,
// "created_at":"2020-04-11T10:06:02Z",
// "updated_at":"2023-09-08T14:02:59Z",
// "private_gists":0,
// "total_private_repos":3,
// "owned_private_repos":3,
// "disk_usage":62232,"collaborators":0,
// "two_factor_authentication":false,
// "plan":{"name":"free",
// "space":976562499,
// "collaborators":0,
// "private_repos":10000}}]]

// map[{"message":"Bad credentials","documentation_url":"https://docs.github.com/rest"}:[]]
