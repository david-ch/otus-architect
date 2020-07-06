package user

import (
	"log"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/user-service/util"
	"github.com/gorilla/mux"
)

type getRs struct {
	ID             int64  `json:"id"`
	Username       string `json:"userName"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	PersonalStatus string `json:"personalStatus"`
}

func onGet(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Header.Get("X-Username")
	username := mux.Vars(r)["username"]

	log.Println("GET: currentUser = " + currentUser + " username = " + username)

	if currentUser != username && currentUser != "tech" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u, err := users.findByUsername(username)

	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	rs := getRs{
		ID:             u.ID,
		Username:       u.Username,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		PersonalStatus: u.PersonalStatus,
	}

	util.Resp(w, http.StatusOK, rs)
}
