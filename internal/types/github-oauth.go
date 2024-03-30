package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GithubAccessTokenResponse struct {
	Error       string `json:"error"`
	Error_desc  string `json:"error_description"`
	AccessToken string `json:"access_token"`
}
type GithubDataUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Avatar_url string `json:"avatar_url"`
}
type GitHubMailStruct struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

var gitScopes = "user user:email"

func GitOatuhStr(state string) string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=%v&scope=%v&state=%v", os.Getenv("GH_CLIENT_ID"), os.Getenv("GITHUB_OAUTH_CALLBACK"), gitScopes, state)
}

func GetGitHubPayload(bearer string) (*GithubDataUser, *GitHubMailStruct, error) {
	r, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	rEmail, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	//
	r.Header.Set("Authorization", bearer)
	rEmail.Header.Set("Authorization", bearer)
	//
	response, err := http.DefaultClient.Do(r)
	responseE, errE := http.DefaultClient.Do(rEmail)
	defer response.Body.Close()
	defer responseE.Body.Close()
	if err != nil || errE != nil {
		return &GithubDataUser{}, &GitHubMailStruct{}, err
	}
	var (
		payload GithubDataUser
		mails   []GitHubMailStruct
	)

	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return &GithubDataUser{}, &GitHubMailStruct{}, err
	}
	if err := json.NewDecoder(responseE.Body).Decode(&mails); err != nil {
		return &GithubDataUser{}, &GitHubMailStruct{}, err
	}
	return &payload, &mails[0], nil
}

func GetGithubAccessToken(code string) (*GithubAccessTokenResponse, error) {
	clientID := os.Getenv("GH_CLIENT_ID")
	clientSecret := os.Getenv("GH_CLIENT_SCRT")
	str := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%v&client_secret=%v&code=%v&", clientID, clientSecret, code)
	req, _ := http.NewRequest("POST", str, nil)
	req.Header.Set("Accept", "application/json")
	var gitResponse GithubAccessTokenResponse
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		return &gitResponse, resperr
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&gitResponse); err != nil {
		return &gitResponse, resperr
	}
	if gitResponse.Error != "" {
		return &gitResponse, fmt.Errorf(gitResponse.Error_desc)
	}
	return &gitResponse, nil
}
