package auth

import (
	"net/http"

	"github.com/david-ch/otus-architect/auth-service/metrics"
	"github.com/gorilla/mux"
)

// RegisterAuthRoutes registers all the routes of auth package
func RegisterAuthRoutes(r *mux.Router) {
	regPath(r, "/register", "/register", "POST", onRegister)
	regPath(r, "/login", "/login", "POST", onLogin)
	regPath(r, "/auth", "/auth", "GET", onAuth)
	regPath(r, "/logout", "/logout", "POST", onLogout)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
