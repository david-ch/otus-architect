package main

import (
	"github.com/david-ch/otus-architect/order-service/health"
	"github.com/david-ch/otus-architect/order-service/metrics"
	"github.com/david-ch/otus-architect/order-service/order"
	"github.com/david-ch/otus-architect/order-service/order/event"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	health.RegisterHealthRoutes(r)
	metrics.RegisterMetricsRoutes(r)
	order.RegisterOrderRoutes(r)
	event.Register()

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
