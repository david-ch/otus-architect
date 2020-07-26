package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/david-ch/otus-architect/user-service/password"
	"github.com/david-ch/otus-architect/user-service/util"
	"github.com/gorilla/mux"
)

type putRq struct {
	Username       string `json:"username"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Password       string `json:"password"`
	PersonalStatus string `json:"personalStatus"`
}

func onPut(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Header.Get("X-Username")
	username := mux.Vars(r)["username"]

	log.Println("PUT: currentUser = " + currentUser + " username = " + username)

	if currentUser != username && currentUser != "tech" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rq := &putRq{}
	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Cannot parse request"})
		return
	}

	passwordHash, err := password.HashPwd(rq.Password)
	if err != nil {
		util.Resp(w, http.StatusInternalServerError, util.ErrorResp{Message: "Error while calculating password hash"})
		return
	}

	u := &user{
		Username:       rq.Username,
		FirstName:      rq.FirstName,
		LastName:       rq.LastName,
		PasswordHash:   passwordHash,
		PersonalStatus: rq.PersonalStatus,
	}

	err = users.updateByUsername(username, u)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
