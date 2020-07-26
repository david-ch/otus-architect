package user

import (
	"net/http"

	"github.com/david-ch/otus-architect/user-service/metrics"
	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers all the routes of user package
func RegisterUserRoutes(r *mux.Router) {
	regPath(r, "/profile/{username}", "/profile", "GET", onGet)
	regPath(r, "/profile", "/profile", "POST", onPost)
	regPath(r, "/profile/{username}", "/profile", "PUT", onPut)
	regPath(r, "/profile/{username}", "/profile", "DELETE", onDelete)
	regPath(r, "/validateCreds", "/validateCreds", "POST", onValidateCreds)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
