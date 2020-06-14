package product

import (
	"github.com/david-ch/otus-architect/hw15/product-service/product/controller"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/product-service/metrics"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers all the routes of product package
func RegisterProductRoutes(r *mux.Router) {
	regPath(r, "/{id}", "/", "GET", controller.OnGet)
	regPath(r, "/", "/", "POST", controller.OnPost)
	regPath(r, "/{id}", "/", "PUT", controller.OnPut)
	regPath(r, "/{id}", "/", "DELETE", controller.OnDelete)
}

func regPath(r *mux.Router, path string, metricsName string, httpMethod string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(path, metrics.WithMetrics(metricsName, handler)).Methods(httpMethod)
}
