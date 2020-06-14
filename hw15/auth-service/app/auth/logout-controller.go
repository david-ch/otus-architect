package auth

import (
	"net/http"
	"time"

	"github.com/david-ch/otus-architect/hw15/auth-service/session"
)

func onLogout(w http.ResponseWriter, r *http.Request) {
	sessionIDCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session.RemoveSession(sessionIDCookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:    "sessionId",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}
