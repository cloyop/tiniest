package middleware

import (
	"net/http"

	"github.com/cloyop/tiniest/internal/types"
)

func Logged(hf func(w http.ResponseWriter, r *http.Request, s types.Session)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := types.GetSession(r)
		if !s.Valid {
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Add("HX-Redirect", "/sign")
				w.WriteHeader(http.StatusForbidden)
				return
			} else {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
		}
		hf(w, r, s)
	}

}
