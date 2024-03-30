package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/helpers"
	"github.com/cloyop/tiniest/internal/templates/pages"
	"github.com/cloyop/tiniest/internal/types"
	"github.com/cloyop/tiniest/internal/utils"
	"github.com/gorilla/mux"
)

func response(w http.ResponseWriter, r *http.Request, execs templ.Component, status int) {
	w.WriteHeader(status)
	if execs != nil {
		if err := execs.Render(r.Context(), w); err != nil {
			fmt.Println(err)
			return
		}
	}
}
func feedBack(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fb := &types.FeedBack{}
		fb.Name = r.FormValue("name")
		fb.Email = r.FormValue("email")
		fb.Subject = r.FormValue("subject")
		fb.Description = r.FormValue("description")
		if errs, hasErrs := fb.Validate(); hasErrs {
			response(w, r, helpers.RequestErrors(errs), http.StatusBadRequest)
			return
		}
		fb.Date = strings.Split(time.Now().String(), ".")[0]
		if err := database.InsertFeedBack(*fb); err != nil {
			response(w, r, helpers.SingleError("Internal error try again"), http.StatusInternalServerError)
			return
		}
		response(w, r, nil, http.StatusOK)
		return
	}
	response(w, r, nil, http.StatusMethodNotAllowed)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	qr := r.URL.Query().Get("qr")
	pair, exists, err := database.GetPair(key)
	if err != nil {
		fmt.Println("error getting pair", key)
		return
	}
	if !exists {
		fmt.Fprintln(w, "Cant Found This Route")
		return
	}
	switch r.Method {
	case "GET":
		if pair.Protected {
			ext := key
			if qr != "" {
				ext += "?qr=true"
			}
			response(w, r, pages.ProtectedRoute(ext), http.StatusOK)
			return
		}
		if pair.User_id != "" {
			utils.PassToIpChan(func() {
				visit, err := types.NewVisit(pair, qr, r)
				if err != nil {
					fmt.Println(err, "on visit")
					return
				}
				if err := database.InsertVisit(visit); err != nil {
					fmt.Println(err, "Inserting")
				}
			})

		}
		http.Redirect(w, r, pair.Long, http.StatusSeeOther)
		return
	case "POST":
		r.ParseForm()
		pwd := r.FormValue("password")
		if pair.Password != pwd {
			response(w, r, helpers.SingleError("password unmatch"), http.StatusBadRequest)
			return
		}
		utils.PassToIpChan(func() {
			visit, err := types.NewVisit(pair, qr, r)
			if err != nil {
				fmt.Println(err, "on visit")
				return
			}
			if err := database.InsertVisit(visit); err != nil {
				fmt.Println(err, "Inserting")
			}
		})

		w.Header().Add("HX-Redirect", pair.Long)
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
