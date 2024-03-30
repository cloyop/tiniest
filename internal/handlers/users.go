package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/helpers"
	"github.com/cloyop/tiniest/internal/templates/partials"
	"github.com/cloyop/tiniest/internal/types"
	"github.com/cloyop/tiniest/internal/utils"
)

func signUpUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if s := types.GetSession(r); s.Valid {
			w.Header().Add("HX-Request", "/dashboard")
			w.WriteHeader(http.StatusOK)
			return
		}
		user := &types.User{}
		if errors, hasErrors := user.Validate(r, false); hasErrors {
			response(w, r, helpers.RequestErrors(*errors), http.StatusBadRequest)
			return
		}
		_, exist, err := database.GetUserByMail(user.Email)
		if err != nil {
			response(w, r, helpers.SingleError("Internal Error, please try again later"), http.StatusInternalServerError)
			return
		}
		if exist {
			response(w, r, helpers.SingleError("Email on use"), http.StatusBadRequest)
			return
		}
		user.GenerateID()
		response(w, r, helpers.Email_Sign_Verification_Form(user.Id), http.StatusOK)
		hash, _ := utils.HashPwd(user.Password)
		user.Password = hash
		user.Protected = true
		user.Joined()
		user.DefaultAvatar()
		utv := utils.NewUserToVerify(user)
		utils.SendingMail(user.Email, "sign up verification", helpers.CodeVerificationEmail_SignUp(utv.Code, utv.User.Name))
		return
	}
}
func signInUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if s := types.GetSession(r); s.Valid {
			w.Header().Add("HX-Redirect", "/dashboard")
			w.WriteHeader(http.StatusOK)
			return
		}
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
		dbUser, exist, err := database.GetUserByMail(email)
		if err != nil {
			response(w, r, helpers.SingleError("Internal Error, please try again later"), http.StatusInternalServerError)
			return
		}
		if !exist || !dbUser.Protected || utils.CompPwd(dbUser.Password, password) != nil {
			response(w, r, helpers.SingleError("Invalid Credentials"), http.StatusBadRequest)
			return
		}
		http.SetCookie(w, types.NewSessionCookie(types.NewSession(dbUser)))
		w.Header().Add("HX-Redirect", "/dashboard")
		w.WriteHeader(http.StatusOK)
		return
	}
}
func signOutUser(w http.ResponseWriter, r *http.Request, s types.Session) {
	if r.Method == "POST" {
		s.RemoveSession()
		http.SetCookie(w, types.NewSessionCookie(""))
		w.Header().Add("HX-Redirect", "/sign")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func updateUser(w http.ResponseWriter, r *http.Request, s types.Session) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		var emailV bool
		var passwordV bool
		var errors []string
		date := time.Now().Unix()
		if r.FormValue("v") != "" {
			currentPassword := r.FormValue("current-password")
			if utils.CompPwd(s.User.Password, currentPassword) != nil {
				response(w, r, helpers.SingleError("Invalid Password"), http.StatusBadRequest)
				return
			}
			code := r.FormValue("code")
			u, err := utils.ValidateUserVerifyCode(s.User.Id, code)
			if err != nil {
				response(w, r, helpers.SingleError("Invalid Code"), http.StatusBadRequest)
				return
			}
			s.User = *u
		}
		if pwd := r.FormValue("password"); pwd != "" {
			emailcut := strings.Split(s.User.Email, "@")
			if len(pwd) > 8 && len(pwd) < 70 && !strings.Contains(pwd, emailcut[0]) && !strings.Contains(pwd, emailcut[1]) && utils.CompPwd(s.User.Password, pwd) != nil {
				hash, _ := utils.HashPwd(pwd)
				s.User.Password = hash
				passwordV = true
			} else {
				errors = append(errors, "Invalid Password")
			}
		}
		if name := r.FormValue("name"); name != "" {
			if len(name) > 4 && len(name) < 30 {
				s.User.Name = name
				s.User.Updates.Name_Last_Update = date
			} else {
				errors = append(errors, "Invalid Name")
			}
		}
		if email := r.FormValue("email"); email != "" {
			if !passwordV && !s.User.Protected {
				errors = append(errors, "Need to set a password first")
			} else {
				if types.EmailPattern.Match([]byte(email)) && email != s.User.Email {
					_, is, err := database.GetUserByMail(email)
					if err != nil {
						response(w, r, helpers.SingleError("Internal Error, Try again"), http.StatusInternalServerError)
						return
					}
					if is {
						errors = append(errors, "Email on use")
					}
					s.User.Email = email
					emailV = true
					s.User.Updates.Email_Last_Update = date
				} else {
					errors = append(errors, "Invalid Email")
				}
			}
		}
		if len(errors) > 0 {
			response(w, r, helpers.RequestErrors(errors), http.StatusBadRequest)
			return
		}
		if emailV || passwordV && s.User.Protected {
			if passwordV {
				s.User.Updates.Password_Last_Update = date
			}
			utv := utils.NewUserToVerify(&s.User)
			utils.SendingMail(s.User.Email, "Verication email code to update your information", helpers.CodeVerificationEmail_UpdateUser(utv.Code, utv.User.Name))
			response(w, r, helpers.Update_Verification_Form(), http.StatusOK)
			return
		}
		s.User.Protected = true
		if err := database.UpdateUser(&s.User); err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError("Internal Error, try again"), http.StatusInternalServerError)
			return
		}
		s.User.Updates.FullFill()
		s.Update()
		w.Header().Add("HX-Refresh", "true")
		w.WriteHeader(http.StatusOK)
		return
	case "GET":
		switch r.URL.Query().Get("temp") {
		case "uu":
			utv, err := utils.UpdateCode(s.User.Id)
			if err != nil {
				response(w, r, helpers.SingleError(err.Error()), http.StatusBadRequest)
				return
			}
			utils.SendingMail(utv.User.Email, "Verication email code to update your information", helpers.CodeVerificationEmail_UpdateUser(utv.Code, utv.User.Name))
			response(w, r, helpers.User_Update_Counter(), http.StatusOK)
			return
		case "pi":
			response(w, r, partials.PersonalInfo_Pre(&s), http.StatusOK)
			return
		case "pf":
			response(w, r, partials.PersonalInfo_Form(&s), http.StatusOK)
			return
		}
	case "DELETE":
		ps, _ := database.GetAllPairsByID(s.User.Id)

		if err := database.DeleteAllViewFromUser(s.User.Id); err != nil {
			fmt.Println("trying to delete all views from", s.User.Email)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if errPairs := database.DeleteAllPairFromUser(s.User.Id); errPairs != nil {
			fmt.Println("trying to delete all views from ", s.User.Email, "->", errPairs)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		go func() {
			for _, v := range *ps {
				v.DeleteQR()
			}
		}()
		if err := database.DeleteUser(s.User.Email); err != nil {
			fmt.Println(err, "trying to delete User", s.User.Email)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s.RemoveSession()
		http.SetCookie(w, types.NewSessionCookie(""))
		w.Header().Add("HX-Redirect", "/sign")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func forgotPassword(w http.ResponseWriter, r *http.Request) {
	if s := types.GetSession(r); s.Valid {
		w.Header().Add("hx-Redirect", "/dashboard/account")
		w.WriteHeader(http.StatusOK)
		return
	}
	r.ParseForm()
	uid := r.FormValue("u")

	switch r.Method {
	case "PUT":
		utv, err := utils.UpdateCode(uid)
		if err != nil {
			w.Header().Add("HX-Redirect", "/sign")
			response(w, r, helpers.SingleError("Something went wrong, please issue your password recovery again"), http.StatusInternalServerError)
			return
		}
		utils.SendingMail(utv.User.Email, "Password recovery code", helpers.CodeVerificationEmail_PasswordRecovery(utv.Code))
		response(w, r, helpers.ForgotPassword_Counter(uid), http.StatusOK)
		return
	case "GET":
		if r.FormValue("g") != "" {
			response(w, r, helpers.ForgotPassword_E_Getter(), http.StatusOK)
			return
		}
		email := r.FormValue("email")
		if email == "" || !types.EmailPattern.Match([]byte(email)) {
			response(w, r, helpers.SingleError("Invalid Email"), http.StatusBadRequest)
			return
		}
		u, is, err := database.GetUserByMail(email)
		if err != nil {
			response(w, r, helpers.SingleError("Internal Error,Try Again"), http.StatusInternalServerError)
			return
		}
		if !is {
			response(w, r, helpers.SingleError("email do not exist"), http.StatusBadRequest)
			return
		}
		utv := utils.NewUserToVerify(u)
		utils.SendingMail(email, "Password recovery code", helpers.CodeVerificationEmail_PasswordRecovery(utv.Code))
		response(w, r, helpers.Email_FP_Verification_Form(u.Id), http.StatusOK)
		return
	case "POST":
		code := r.FormValue("code")
		u, err := utils.ValidateUserVerifyCode(uid, code)
		if err != nil {
			response(w, r, helpers.SingleError(err.Error()), http.StatusInternalServerError)
			return
		}
		newPwd := r.FormValue("new-password")
		if len(newPwd) < 10 || len(newPwd) > 72 {
			response(w, r, helpers.SingleError("Invalid Password"), http.StatusBadRequest)
			return
		}
		u.Password, err = utils.HashPwd(newPwd)
		if err != nil {
			response(w, r, helpers.SingleError(err.Error()), http.StatusInternalServerError)
			return
		}
		if err := database.UpdatePassword(u.Password, u.Id); err != nil {
			response(w, r, helpers.SingleError("Internal Error, try again"), http.StatusInternalServerError)
			return
		}
		u.Protected = true
		http.SetCookie(w, types.NewSessionCookie(types.NewSession(u)))
		w.Header().Add("HX-Redirect", "/dashboard")
		w.WriteHeader(http.StatusOK)
		return
	}
}
