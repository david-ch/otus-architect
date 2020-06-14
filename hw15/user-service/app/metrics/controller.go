package metrics

import (
	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterMetricsRoutes registers all the routes of metrics package
func RegisterMetricsRoutes(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler())
	registerMetrics()
}
