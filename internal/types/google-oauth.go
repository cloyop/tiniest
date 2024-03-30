package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GoogleProfileDataResponse struct {
	Id       string      `json:"id"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Verified bool        `json:"verified_email"`
	Picture  string      `json:"picture"`
	Locale   string      `json:"es"`
	Error    GoogleError `json:"error"`
}
type GoogleError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GoogleAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Error       string `json:"error"`
	Error_desc  string `json:"error_description"`
}

var googleScopes = "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"

func GoogleOauthStr(state string) string {
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?scope=%v&access_type=offline&include_granted_scopes=true&response_type=code&state=%v&redirect_uri=%v&client_id=%v", googleScopes, state, os.Getenv("GOOGLE_OAUTH_CALLBACK"), os.Getenv("GOOGLE_CLIENT_ID"))
}
func GetGoogleAccessToken(code string) (*GoogleAccessTokenResponse, error) {
	s := fmt.Sprintf("https://oauth2.googleapis.com/token?code=%v&client_id=%v&client_secret=%v&grant_type=%v&redirect_uri=%v", code, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SCRT"), "authorization_code", os.Getenv("GOOGLE_OAUTH_CALLBACK"))
	req, _ := http.NewRequest("POST", s, nil)
	response, err := http.DefaultClient.Do(req)
	var googleOAuth GoogleAccessTokenResponse
	if err != nil {
		return &googleOAuth, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&googleOAuth); err != nil {
		return &googleOAuth, err
	}
	if googleOAuth.Error != "" {
		return &googleOAuth, fmt.Errorf(googleOAuth.Error_desc + " " + googleOAuth.Error)
	}
	return &googleOAuth, nil
}
func GetGoogleUserProfile(accessToken string) (*GoogleProfileDataResponse, error) {
	userProfile := &GoogleProfileDataResponse{}
	request, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	responseUserData, err := http.DefaultClient.Do(request)
	if err != nil {
		return userProfile, err
	}
	defer responseUserData.Body.Close()
	if err := json.NewDecoder(responseUserData.Body).Decode(&userProfile); err != nil {
		return userProfile, err
	}
	return userProfile, nil
}
