package warehouse

import (
	"github.com/david-ch/otus-architect/warehouse-service/metrics"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWarehouseRoutes registers all the routes of payment package
func RegisterWarehouseRoutes(r *mux.Router) {
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
