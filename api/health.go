package api

import (
	"net/http"
)

// HealthHandler returns a http handler that checks service status
func HealthHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}
}
