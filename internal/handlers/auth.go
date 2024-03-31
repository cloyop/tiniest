package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/helpers"
	"github.com/cloyop/tiniest/internal/templates/partials"
	"github.com/cloyop/tiniest/internal/types"
	"github.com/cloyop/tiniest/internal/utils"
	"github.com/gorilla/mux"
)

func gitHubOAuth(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	code := queries.Get("code")
	errQ := queries.Get("error")
	if errQ != "" {
		errSTR := queries.Get("error_description")
		fmt.Println(errSTR)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	if queries.Get("state") != os.Getenv("OAUTH_STATE") || code == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	gtr, err := types.GetGithubAccessToken(code)
	if err != nil {
		fmt.Println("Getting Access Token github", err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	github_User, github_Mail, err := types.GetGitHubPayload("Bearer " + gtr.AccessToken)
	if err != nil {
		fmt.Println("Getting User Profile github", err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}

	u, exist, err := database.GetUserByMail(github_Mail.Email)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	if !exist {
		u.GenerateID()
		u.Joined()
		u.Email = github_Mail.Email
		u.Name = github_User.Name
		u.Avatar_url = github_User.Avatar_url
		u.Verified = github_Mail.Verified
		if err := database.InsertUser(u); err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/sign", http.StatusSeeOther)
			return
		}
		utils.MigratePublicToNewUser(*u, w, r)
	}
	http.SetCookie(w, types.NewSessionCookie(types.NewSession(u)))
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
func googleOAuth(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	code := queries.Get("code")
	errQ := queries.Get("error")
	if errQ != "" {
		errSTR := queries.Get("error_description")
		fmt.Println(errSTR)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	if queries.Get("state") != os.Getenv("OAUTH_STATE") || code == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	googleOAuth, err := types.GetGoogleAccessToken(code)
	if err != nil {
		fmt.Println("Getting Access Token google", err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	userProfile, err := types.GetGoogleUserProfile(googleOAuth.AccessToken)
	if err != nil {
		fmt.Println("Getting User Profile google", err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	u, exist, err := database.GetUserByMail(userProfile.Email)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	}
	if !exist {
		u.GenerateID()
		u.Email = userProfile.Email
		u.Name = userProfile.Name
		u.Joined()
		u.Avatar_url = userProfile.Picture
		u.Verified = userProfile.Verified
		if err := database.InsertUser(u); err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/sign", http.StatusSeeOther)
			return
		}
		utils.MigratePublicToNewUser(*u, w, r)
	}
	http.SetCookie(w, types.NewSessionCookie(types.NewSession(u)))
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
func signForm(w http.ResponseWriter, r *http.Request) {
	if s := types.GetSession(r); s.Valid {
		w.Header().Add("HX-Redirect", "/dashboard")
		w.WriteHeader(204)
		return
	}
	switch mux.Vars(r)["to"] {
	case "oauth":
		OAM := r.URL.Query().Get("m")
		if OAM == "git" {
			s := types.GitOatuhStr(os.Getenv("OAUTH_STATE"))
			http.Redirect(w, r, s, http.StatusSeeOther)
			return
		}
		if OAM == "goo" {
			s := types.GoogleOauthStr(os.Getenv("OAUTH_STATE"))
			http.Redirect(w, r, s, http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/sign", http.StatusSeeOther)
		return
	case "up":
		response(w, r, partials.SignForm(false, "up"), http.StatusOK)
		return
	case "in":
		response(w, r, partials.SignForm(true, "in"), http.StatusOK)
		return
	}
}
func tiniestAuth(w http.ResponseWriter, r *http.Request) {
	uidKey := mux.Vars(r)["uidkey"]
	switch r.Method {
	case "GET":
		utv, err := utils.UpdateCode(uidKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response(w, r, helpers.SingleError(err.Error()), http.StatusBadRequest)
			return
		}
		utils.SendingMail(utv.User.Email, "sign up verification", helpers.CodeVerificationEmail_SignUp(utv.Code, utv.User.Name))
		response(w, r, helpers.Email_Sign_Counter(uidKey), http.StatusOK)
		return
	case "POST":
		if s := types.GetSession(r); s.Valid {
			w.Header().Add("HX-Redirect", "/dashboard")
			w.WriteHeader(http.StatusOK)
			return
		}
		r.ParseForm()
		code := r.FormValue("code")
		user, err := utils.ValidateUserVerifyCode(uidKey, code)
		if err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError(err.Error()), http.StatusBadRequest)
			return
		}
		user.Verified = true
		if err := database.InsertUser(user); err != nil {
			response(w, r, helpers.SingleError(err.Error()), http.StatusBadRequest)
			return
		}
		utils.MigratePublicToNewUser(*user, w, r)
		http.SetCookie(w, types.NewSessionCookie(types.NewSession(user)))
		w.Header().Add("HX-Redirect", "/dashboard")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
