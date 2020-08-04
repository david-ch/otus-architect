package user

import (
	"net/http"

	"github.com/david-ch/otus-architect/user-service/util"
	"github.com/gorilla/mux"
)

func onDelete(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Header.Get("X-Username")
	username := mux.Vars(r)["username"]

	if currentUser != username && currentUser != "tech" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := users.deleteByUsername(username)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
