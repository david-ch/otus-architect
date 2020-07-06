package auth

import (
	"encoding/json"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/auth-service/client"
	"github.com/david-ch/otus-architect/hw15/auth-service/session"
	"github.com/david-ch/otus-architect/hw15/auth-service/util"
)

type loginRq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func onLogin(w http.ResponseWriter, r *http.Request) {
	rq := loginRq{}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	isValid, err := client.ValidateCreds(rq.Username, rq.Password)
	if err != nil {
		util.Resp(w, http.StatusInternalServerError, util.FromError(err))
		return
	}

	if isValid {
		sess := session.CreateSession(rq.Username)

		http.SetCookie(w, &http.Cookie{
			Name:  "sessionId",
			Value: sess.ID,
			Path:  "/",
		})

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
