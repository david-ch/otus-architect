package delivery

import (
	"github.com/david-ch/otus-architect/delivery-service/metrics"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterDeliveryRoutes registers all the routes of payment package
func RegisterDeliveryRoutes(r *mux.Router) {
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
