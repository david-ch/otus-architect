package order

import (
	"github.com/david-ch/otus-architect/order-service/metrics"
	"github.com/david-ch/otus-architect/order-service/order/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterOrderRoutes registers all the routes of product package
func RegisterOrderRoutes(r *mux.Router) {
	regPath(r, "/{id}", "/", "GET", controller.OnGet)
	regPath(r, "/", "/", "POST", controller.OnPost)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
