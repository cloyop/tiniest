package handlers

import (
	"net/http"

	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/pages"
	"github.com/cloyop/tiniest/internal/types"
	"github.com/cloyop/tiniest/internal/utils"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var pairs []types.PairLinks
		pairsStr := utils.GetPairCookieString(r)
		if pairsStr != "" {
			pairs, pairsStr = utils.GetPairsFromCookie(utils.ProcessCookiesString(pairsStr))
			http.SetCookie(w, utils.NewPublicPairCookie("", pairsStr))
		}
		s := types.GetSession(r)
		response(w, r, pages.Home(&s, pairs), http.StatusOK)
		return
	}
}
func dashboard(w http.ResponseWriter, r *http.Request, s types.Session) {
	if r.Method == "GET" {
		pairs, err := database.GetAllPairsByID(s.User.Id)
		if err == nil && s.User.LinksCount != len(*pairs) {
			s.User.LinksCount = len(*pairs)
			if err := database.UpdateUserLinksCount(&s.User); err == nil {
				s.Update()
			}
		}
		response(w, r, pages.Dashboard(&s, *pairs), http.StatusOK)
		return
	}
	http.Redirect(w, r, "/sign", http.StatusSeeOther)
}
func dashboardAccount(w http.ResponseWriter, r *http.Request, s types.Session) {
	if r.Method == "GET" {
		response(w, r, pages.DashBoard_Account(&s), http.StatusOK)
		return
	}
}
func sign(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if s := types.GetSession(r); s.Valid {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		query := r.URL.Query().Get("m")
		if query == "up" {
			response(w, r, pages.Sign(false, "up"), http.StatusOK)
			return
		}
		response(w, r, pages.Sign(true, "in"), http.StatusOK)
		return
	}
}
func termsOfService(w http.ResponseWriter, r *http.Request) {
	pages.TermsOfService().Render(r.Context(), w)
}
func privacyPolicy(w http.ResponseWriter, r *http.Request) {
	pages.PrivacyPolicy().Render(r.Context(), w)
}
