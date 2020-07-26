package health

import (
	"github.com/david-ch/otus-architect/delivery-service/util"
	"net/http"
)

type health struct {
	Status string
}

func onHealth(w http.ResponseWriter, r *http.Request) {
	util.Resp(w, http.StatusOK, health{"OK"})
}