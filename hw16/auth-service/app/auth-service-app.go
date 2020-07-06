package main

import (
	"log"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/auth-service/auth"
	"github.com/david-ch/otus-architect/hw15/auth-service/health"
	"github.com/david-ch/otus-architect/hw15/auth-service/metrics"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	metrics.RegisterMetricsRoutes(r)
	health.RegisterHealthRoutes(r)
	auth.RegisterAuthRoutes(r)

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
