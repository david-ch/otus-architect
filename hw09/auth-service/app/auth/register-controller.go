package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/david-ch/otus-architect/hw09/auth-service/client"
	"github.com/david-ch/otus-architect/hw09/auth-service/util"
)

type registerRq struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

func onRegister(w http.ResponseWriter, r *http.Request) {
	rq := registerRq{}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Cannot parse request"})
		return
	}

	if !validateRegisterRq(rq) {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Bad request format"})
		return
	}

	client.CreateUser(rq.Username, rq.FirstName, rq.LastName, rq.Password)

	w.WriteHeader(http.StatusCreated)
}

func validateRegisterRq(rq registerRq) bool {
	return strings.TrimSpace(rq.Username) != "" &&
		strings.TrimSpace(rq.FirstName) != "" &&
		strings.TrimSpace(rq.LastName) != "" &&
		strings.TrimSpace(rq.Password) != ""
}
