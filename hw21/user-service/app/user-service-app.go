package main

import (
	"github.com/david-ch/otus-architect/user-service/health"
	"github.com/david-ch/otus-architect/user-service/metrics"
	"log"
	"net/http"

	"github.com/david-ch/otus-architect/user-service/user"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	metrics.RegisterMetricsRoutes(r)
	health.RegisterHealthRoutes(r)
	user.RegisterUserRoutes(r)

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
