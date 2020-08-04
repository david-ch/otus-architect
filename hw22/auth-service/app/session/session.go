package session

import (
	"sync"

	"github.com/google/uuid"
)

// Session describes data that is stored in the session
type Session struct {
	ID       string
	Username string
}

type storage struct {
	idx map[string]Session
	mux *sync.Mutex // todo replace with RW lock
}

var stor = &storage{
	idx: make(map[string]Session, 0),
	mux: &sync.Mutex{},
}

// CreateSession creates new session for specified user
func CreateSession(username string) Session {
	stor.mux.Lock()
	defer stor.mux.Unlock()

	session := Session{
		ID:       uuid.New().String(),
		Username: username,
	}

	// todo remove all the other sessions of the user
	stor.idx[session.ID] = session

	return session
}

// FindSession looks for session with the id. Returns false as second parameter if session not found.
func FindSession(id string) (s Session, ok bool) {
	stor.mux.Lock()
	defer stor.mux.Unlock()

	s, ok = stor.idx[id]
	return
}

// RemoveSession removes session from storage
func RemoveSession(id string) {
	stor.mux.Lock()
	defer stor.mux.Unlock()

	delete(stor.idx, id)
}
