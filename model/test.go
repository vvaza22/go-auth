package model

import (
	"auth/account"
	"auth/database"
	"auth/session"
	"fmt"
	"net/http"
)

func TestPageHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	ac := account.NewAccountManager(dbSource)
	ok := sm.UserLoggedIn()
	if ok {
		fmt.Fprintln(w, "YES")
		fmt.Fprintf(w, "%s\n", sm.UserAccount().FirstName)
	} else {
		fmt.Fprintln(w, "NO")
	}
	a, _ := ac.GetAccount("vazzu")
	//sess.Set("user", a)
	sm.SetUser(a)
}
