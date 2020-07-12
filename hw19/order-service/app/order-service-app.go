package main

import (
	"github.com/david-ch/otus-architect/hw19/order-service/health"
	"github.com/david-ch/otus-architect/hw19/order-service/metrics"
	"github.com/david-ch/otus-architect/hw19/order-service/order"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	health.RegisterHealthRoutes(r)
	metrics.RegisterMetricsRoutes(r)
	order.RegisterProductRoutes(r)

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
