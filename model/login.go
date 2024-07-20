package model

import (
	"auth/account"
	"auth/database"
	"auth/page"
	"auth/session"
	"auth/utility"
	"encoding/json"
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	var pageData *page.Data = page.NewData("Login", sm)
	var t *template.Template = utility.ParseTemplates("layout", "login")
	_ = t.ExecuteTemplate(w, "layout", *pageData)
}

type LoginResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	w.Header().Set("Content-Type", "application/json")
	response := LoginResponse{Status: false, Message: "Invalid username or password"}

	r.ParseForm()

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// Get the account manager
	am := account.NewAccountManager(dbSource)

	if am.AccountExists(username) {
		account, _ := am.GetAccount(username)
		if account.CheckPassword(password) {
			response.Status = true
			response.Message = "Success"
			// Update the session
			sm.SetUser(account)
		}
	}

	// Return the response to the client
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}
