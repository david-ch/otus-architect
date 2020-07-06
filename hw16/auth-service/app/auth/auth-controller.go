package auth

import (
	"net/http"

	"github.com/david-ch/otus-architect/hw15/auth-service/session"
)

func onAuth(w http.ResponseWriter, r *http.Request) {
	sessionIDCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sess, ok := session.FindSession(sessionIDCookie.Value)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Add("X-Username", sess.Username)
	w.WriteHeader(http.StatusOK)
}
