package main

import (
	"net/http"
	"os"

	"github.com/cloyop/tiniest/internal/config"
	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnvConfig()
	config.LoadHelpers()
	defer database.ConnectDB().Close()
	r := mux.NewRouter()
	handlers.LoadHandlers(r)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
