package user

import (
	"encoding/json"
	"net/http"

	"github.com/david-ch/otus-architect/hw09/user-service/password"
	"github.com/david-ch/otus-architect/hw09/user-service/util"
)

type validateCredsRq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type validateCredsRs struct {
	IsValid bool `json:"isValid"`
}

func onValidateCreds(w http.ResponseWriter, r *http.Request) {
	rq := &validateCredsRq{}

	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Cannot parse request"})
		return
	}

	u, err := users.findByUsername(rq.Username)

	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.ErrorResp{Message: "Error whild searching user"})
		return
	}

	isValid := password.ValidatePwd(u.PasswordHash, rq.Password)

	util.Resp(w, http.StatusOK, validateCredsRs{IsValid: isValid})
}
