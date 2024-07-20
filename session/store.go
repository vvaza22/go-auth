package session

import (
	"auth/utility"
	"sync"
)

const (
	IdLength = 32
)

type Store struct {
	store map[string]*Session
	rw    sync.RWMutex
}

var instance *Store
var once sync.Once

func Instance() *Store {
	once.Do(func() {
		instance = &Store{
			store: make(map[string]*Session),
		}
	})
	return instance
}

// Create creates a new session in the store and returns the id of the session
func (s *Store) Create() string {
	s.rw.Lock()
	defer s.rw.Unlock()

	// Generate an id that does not exist in the store
	for {
		sessionId := utility.GenerateRandomString(IdLength)
		_, exists := s.store[sessionId]
		if !exists {
			s.store[sessionId] = NewSession(sessionId)
			return sessionId
		}
	}
}

// Exists checks whether the given session id exists in the store
func (s *Store) Exists(sessionId string) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	_, exists := s.store[sessionId]
	return exists
}

// Get returns a pointer the session
func (s *Store) Get(sessionId string) *Session {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.store[sessionId]
}

// Delete deletes the session from the store
func (s *Store) Delete(sessionId string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	delete(s.store, sessionId)
}
