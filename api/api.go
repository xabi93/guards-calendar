package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New initializes the api handler with routes exposing the routes handlers.
func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", HealthHandler()).Methods(http.MethodGet)

	return router
}
