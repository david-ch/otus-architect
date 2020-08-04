package user

import (
	"encoding/json"
	"net/http"

	"github.com/david-ch/otus-architect/user-service/password"
	"github.com/david-ch/otus-architect/user-service/util"
)

type postRq struct {
	Username       string  `json:"username"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Password       string  `json:"password"`
	PersonalStatus *string `json:"personalStatus,omitempty"`
}

func onPost(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Header.Get("X-Username")
	if currentUser != "tech" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rq := &postRq{}
	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Cannot parse request"})
		return
	}

	personalStatus := "Hey! I am new user!"
	if rq.PersonalStatus != nil {
		personalStatus = *rq.PersonalStatus
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
		PersonalStatus: personalStatus,
	}

	_, err = users.create(u)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	util.Resp(w, http.StatusOK, rq.Username)
}
