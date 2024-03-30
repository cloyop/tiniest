package handlers

import (
	"fmt"
	"net/http"

	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/helpers"
	"github.com/cloyop/tiniest/internal/templates/partials"
	"github.com/cloyop/tiniest/internal/types"
	"github.com/cloyop/tiniest/internal/utils"
)

func newPublicPair(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		p := &types.PairLinks{}
		if errors, hasErrors := p.Validate(r); hasErrors {
			response(w, r, helpers.RequestErrors(errors), http.StatusBadRequest)
			return
		}
		if err := p.LookUpHost(); err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError("Cant reach that host"), http.StatusBadRequest)
			return
		}
		p.Key = utils.ShortURLGen()
		p.DoShort(r.Host)
		p.CreatedAt()
		if err := database.InsertPair(p); err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError("Internal Error, Try again"), http.StatusInternalServerError)
			return
		}
		cookStr := utils.GetPairCookieString(r)
		if len(*utils.ProcessCookiesString(cookStr)) >= 5 {
			response(w, r, helpers.SingleError("max pairs capped, subscribe to create more"), http.StatusBadRequest)
			return
		}
		http.SetCookie(w, utils.NewPublicPairCookie(cookStr, fmt.Sprintf("%v-", p.Key)))
		response(w, r, partials.UrlItem(p), http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func pair(w http.ResponseWriter, r *http.Request, s types.Session) {
	switch r.Method {
	case "POST":
		p := &types.PairLinks{}
		if errors, hasErrors := p.Validate(r); hasErrors {
			response(w, r, helpers.RequestErrors(errors), http.StatusBadRequest)
			return
		}
		if s.User.LinksCount >= s.User.MaxLinksCount {
			response(w, r, nil, http.StatusNotAcceptable)
			return
		}
		s.User.LinksCount++
		p.User_id = s.User.Id
		p.Key = utils.ShortURLGen()
		p.DoShort(r.Host)
		p.CreatedAt()

		if p.Title == "" {
			p.Title, _ = p.GetTittle()
		}

		if err := p.GenerateQR(); err != nil {
			fmt.Println("Generating Qr", err)
			response(w, r, helpers.SingleError("Internal Error, Try again"), http.StatusInternalServerError)
			return
		}
		if err := database.InsertPair(p); err != nil {
			fmt.Println("Inserting Pair", err)
			response(w, r, helpers.SingleError("Internal Error, Try again"), http.StatusInternalServerError)
			return
		}
		s.Update()
		response(w, r, partials.Link_Item(p), http.StatusOK)
		if err := database.UpdateUserLinksCount(&s.User); err != nil {
			fmt.Println(err, "<<<")
		}
		return
	case "DELETE":
		key := r.FormValue("key")
		s.User.LinksCount--
		if err := database.DeleteAllViewsFromKey(key); err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError("Internal Error, try again"), http.StatusInternalServerError)
			return
		}
		if _, err := database.DeletePair(key, s.User.Id); err != nil {
			fmt.Println(err)
			response(w, r, helpers.SingleError("Internal Error, try again"), http.StatusInternalServerError)
			return
		}
		s.Update()
		if err := database.UpdateUserLinksCount(&s.User); err != nil {
			fmt.Println(err, "<<<<")
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func lockPair(w http.ResponseWriter, r *http.Request, s types.Session) {
	key := r.FormValue("key")
	locked := r.FormValue("lock")
	r.ParseForm()
	switch r.Method {
	case "GET":
		if key == "" {
			response(w, r, helpers.SingleError("Missing Key"), http.StatusBadRequest)
			return
		}
		if locked != "" {
			pwd, err := database.GetPairPassword(key, s.User.Id)
			if err != nil {
				response(w, r, helpers.SingleError("Error Try again"), http.StatusInternalServerError)
				return
			}
			response(w, r, partials.LockerLockedForm(key, pwd), http.StatusOK)
			return
		}
		response(w, r, partials.LockerForm(key), http.StatusOK)
		return
	case "POST":
		pwd := r.FormValue("password")
		if len(pwd) < 10 || len(pwd) > 20 {
			response(w, r, helpers.SingleError("Invalid Password Length"), http.StatusBadRequest)
			return
		}
		pair, err := database.LockPair(pwd, key, s.User.Id)
		if err != nil {
			response(w, r, helpers.SingleError("Internal Server Error"), http.StatusInternalServerError)
			return
		}
		response(w, r, partials.Link_Item(pair), http.StatusOK)
		return
	case "DELETE":
		pairs, err := database.DeletePairPassword(key, s.User.Id)
		if err != nil {
			response(w, r, helpers.SingleError("Internal Server Error"), http.StatusInternalServerError)
			return
		}
		response(w, r, partials.Link_Item(&pairs), http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusForbidden)
}
func pairFilter(w http.ResponseWriter, r *http.Request, s types.Session) {
	r.ParseForm()
	queryext := ""
	key := r.FormValue("key")
	queryext += database.AddFilter("key", key)

	if date := r.FormValue("date-filter"); date != "" {
		queryext += database.AddFilter("date", date)
	}
	if from := r.FormValue("from-filter"); from != "" {
		queryext += database.AddFilter("visit_from", from)
	}
	if qr := r.FormValue("qr-filter"); qr != "" {
		if qr == "true" {
			qr = "1"
		} else {
			qr = "0"
		}
		queryext += database.AddFilter("qr", qr)
	}
	visits, _ := database.GetVisitsFiltered(s.User.Id, queryext)
	if r.Method == "POST" {
		response(w, r, partials.TableTrackingBody(*visits), http.StatusOK)
		return
	}
	dates, from := utils.MapFilters(*visits)
	response(w, r, partials.TrackingPair(key, *visits, dates, from), http.StatusOK)
}
