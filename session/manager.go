package session

import (
	"auth/account"
)

const (
	currentUserKey = "user"
)

type Manager struct {
	session *Session
}

func NewSessionManager(session *Session) *Manager {
	return &Manager{session}
}

func (manager *Manager) UserLoggedIn() bool {
	return manager.session.Exists(currentUserKey)
}

func (manager *Manager) SetUser(account account.Account) {
	manager.session.Set(currentUserKey, account)
}

func (manager *Manager) UserAccount() account.Account {
	if manager.UserLoggedIn() {
		return manager.session.Get(currentUserKey).(account.Account)
	}
	return account.Account{}
}

func (manager *Manager) Logout() {
	manager.session.Delete(currentUserKey)
}
