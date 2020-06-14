package main

import (
	"log"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/product-service/health"
	"github.com/david-ch/otus-architect/hw15/product-service/metrics"
	"github.com/david-ch/otus-architect/hw15/product-service/product"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	product.RegisterProductRoutes(r)
	metrics.RegisterMetricsRoutes(r)
	health.RegisterHealthRoutes(r)

	r.Use(commonMiddleware)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
