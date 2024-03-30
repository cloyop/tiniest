package handlers

import (
	"net/http"

	"github.com/cloyop/tiniest/internal/middleware"
	"github.com/gorilla/mux"
)

func LoadHandlers(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/assets/"))))

	userRoute := r.PathPrefix("/user").Subrouter()
	pairRoute := r.PathPrefix("/pair").Subrouter()
	authRoute := r.PathPrefix("/auth").Subrouter()

	userRoute.HandleFunc("/signin", signInUser)
	userRoute.HandleFunc("/signup", signUpUser)
	userRoute.HandleFunc("/signout", middleware.Logged(signOutUser))
	userRoute.HandleFunc("/update", middleware.Logged(updateUser))
	userRoute.HandleFunc("/forgot-password", forgotPassword)

	pairRoute.HandleFunc("/p", newPublicPair)
	pairRoute.HandleFunc("/lock", middleware.Logged(lockPair))
	pairRoute.HandleFunc("/filter", middleware.Logged(pairFilter))
	pairRoute.HandleFunc("/", middleware.Logged(pair))

	authRoute.HandleFunc("/ttcb/{uidkey}", tiniestAuth)
	authRoute.HandleFunc("/ghcb", gitHubOAuth)
	authRoute.HandleFunc("/ggcb", googleOAuth)

	r.HandleFunc("/sign/{to}", signForm)
	r.HandleFunc("/sign", sign)

	r.HandleFunc("/dashboard/account", middleware.Logged(dashboardAccount))
	r.HandleFunc("/dashboard", middleware.Logged(dashboard))
	r.HandleFunc("/feedback", feedBack)
	r.HandleFunc("/privacy-policy", privacyPolicy)
	r.HandleFunc("/terms-of-service", termsOfService)
	r.HandleFunc("/{key}", redirect)
	r.HandleFunc("/", home)
}
