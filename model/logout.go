package model

import (
	"auth/database"
	"auth/session"
	"net/http"
)

func LogoutPostHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	if sm.UserLoggedIn() {
		sm.Logout()
	}
	http.Redirect(w, r, "/", 302)
}
