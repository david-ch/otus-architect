package health

import (
	"net/http"

	"github.com/david-ch/otus-architect/hw15/user-service/util"
	"github.com/gorilla/mux"
)

// RegisterHealthRoutes registers all the routes of health package
func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("/health", onHealth)
}

type health struct {
	Status string
}

func onHealth(w http.ResponseWriter, r *http.Request) {
	util.Resp(w, http.StatusOK, health{"OK"})
}
