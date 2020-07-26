package product

import (
	"github.com/david-ch/otus-architect/product-service/product/controller"
	"net/http"

	"github.com/david-ch/otus-architect/product-service/metrics"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers all the routes of product package
func RegisterProductRoutes(r *mux.Router) {
	regPath(r, "/promoted", "/promoted", "GET", controller.OnGetPromoted)
	regPath(r, "/{id}", "/", "GET", controller.OnGet)
	regPath(r, "/", "/{search}", "GET", controller.OnSearch).Queries("search", "{search}")
	regPath(r, "/", "/", "POST", controller.OnPost)
	regPath(r, "/{id}", "/", "PUT", controller.OnPut)
	regPath(r, "/{id}", "/", "DELETE", controller.OnDelete)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
