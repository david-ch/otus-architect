package payment

import (
	"github.com/david-ch/otus-architect/payment-service/metrics"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterPaymentRoutes registers all the routes of payment package
func RegisterPaymentRoutes(r *mux.Router) {
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
