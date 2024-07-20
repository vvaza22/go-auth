package session

import (
	"net/http"
	"sync"
)

const (
	sessionCookieName     = "session_id"
	sessionCookieHttpOnly = true
	sessionCookieSecure   = false
)

type Session struct {
	ID   string
	data map[string]any
	rw   sync.RWMutex
}

func NewSession(sessionId string) *Session {
	return &Session{
		ID:   sessionId,
		data: make(map[string]any),
	}
}

func (s *Session) Set(key string, value any) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.data[key] = value
}

func (s *Session) Get(key string) any {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.data[key]
}

func (s *Session) Delete(key string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	delete(s.data, key)
}

func (s *Session) Exists(key string) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	_, ok := s.data[key]
	return ok
}

func setSessionCookie(w http.ResponseWriter, sessionId string) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    sessionId,
		HttpOnly: sessionCookieHttpOnly,
		Secure:   sessionCookieSecure,
	})
}

func GetSession(r *http.Request, w http.ResponseWriter) *Session {
	var sessionId string
	store := Instance()
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		// Create a new session
		sessionId = store.Create()
		setSessionCookie(w, sessionId)
	} else {
		// Read the value from the cookie
		sessionId = cookie.Value
		// Check if the value is valid
		if !store.Exists(sessionId) {
			// Create a new session and reset the cookie
			sessionId = store.Create()
			setSessionCookie(w, sessionId)
		}
	}
	return store.Get(sessionId)
}
