package health

import (
	"github.com/gorilla/mux"
)

// RegisterHealthRoutes registers all the routes of health package
func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("/health", onHealth)
}

