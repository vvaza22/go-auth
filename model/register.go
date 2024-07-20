package model

import (
	"auth/account"
	"auth/database"
	"auth/page"
	"auth/session"
	"auth/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

const (
	firstNameMaxLen = 12
	lastNameMaxLen  = 12
	usernameMaxLen  = 10
	passwordMinLen  = 6
)

type ErrorMessage struct {
	Element string `json:"element"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Status bool           `json:"status"`
	Errors []ErrorMessage `json:"errors"`
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	var pageData *page.Data = page.NewData("Register", sm)
	var t *template.Template = utility.ParseTemplates("layout", "register")
	_ = t.ExecuteTemplate(w, "layout", *pageData)
}

func recordError(response *RegisterResponse, element string, error string) {
	response.Status = false
	response.Errors = append(response.Errors, ErrorMessage{Element: element, Message: error})
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	w.Header().Set("Content-Type", "application/json")

	// Read the form values
	r.ParseForm()

	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var response RegisterResponse = RegisterResponse{Status: true, Errors: []ErrorMessage{}}

	// Get the account manager
	am := account.NewAccountManager(dbSource)

	// Regexp
	namePattern := `^[a-zA-Z]+$`
	nameRegexp := regexp.MustCompile(namePattern)

	usernamePattern := `^[a-zA-Z0-9_]+$`
	usernameRegexp := regexp.MustCompile(usernamePattern)

	// first_name
	if firstName == "" {
		recordError(&response, "first_name", "Is required")
	} else if len(firstName) > firstNameMaxLen {
		recordError(&response, "first_name", fmt.Sprintf("Longer than %d characters", firstNameMaxLen))
	} else if !nameRegexp.MatchString(firstName) {
		recordError(&response, "first_name", "Only alphabetic characters allowed")
	}

	// last_name
	if lastName == "" {
		recordError(&response, "last_name", "Is required")
	} else if len(lastName) > lastNameMaxLen {
		recordError(&response, "last_name", fmt.Sprintf("Longer than %d characters", lastNameMaxLen))
	} else if !nameRegexp.MatchString(lastName) {
		recordError(&response, "last_name", "Only alphabetic characters allowed")
	}

	// username
	if username == "" {
		recordError(&response, "username", "Is required")
	} else if len(username) > usernameMaxLen {
		response.Status = false
		response.Errors = append(response.Errors, ErrorMessage{Element: "username", Message: fmt.Sprintf("Longer than %d characters", usernameMaxLen)})
	} else if !usernameRegexp.MatchString(username) {
		response.Status = false
		response.Errors = append(response.Errors, ErrorMessage{Element: "username", Message: "Only alphanumeric characters and _ allowed"})
	} else if am.AccountExists(username) {
		response.Status = false
		response.Errors = append(response.Errors, ErrorMessage{Element: "username", Message: "Username already registered"})
	}

	if password == "" {
		recordError(&response, "password", "Is required")
	} else if len(password) < passwordMinLen {
		recordError(&response, "password", fmt.Sprintf("Shorter than %d characters", passwordMinLen))
	}

	if response.Status {
		// Every check was passed
		am.AddAccount(firstName, lastName, username, password)
		// Set the current user
		ac, err := am.GetAccount("vazzu")
		if err != nil {
			panic(err)
		}
		sm.SetUser(ac)
	}

	// Return the response to the client
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}
