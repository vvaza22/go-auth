package main

import (
	"auth/database"
	"auth/model"
	"auth/session"
	"flag"
	"fmt"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request, source *database.Source, sm *session.Manager)

func main() {
	var dbHost *string = flag.String("db_host", "localhost", "Database host")
	var dbName *string = flag.String("db_name", "", "Database name")
	var dbUsername *string = flag.String("db_username", "", "Database username")
	var dbPassword *string = flag.String("db_password", "", "Database password")

	flag.Parse()

	if *dbName == "" {
		fmt.Println("Database name is required")
		return
	}

	if *dbUsername == "" {
		fmt.Println("Database username is required")
		return
	}

	if *dbPassword == "" {
		fmt.Println("Database password is required")
		return
	}

	// Create a data source
	var dbSource *database.Source = database.NewSource(*dbHost, *dbUsername, *dbPassword, *dbName)

	startServer(dbSource)
}

func startServer(dbSource *database.Source) {
	handlerWrapper := func(handler RouteHandler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var sm *session.Manager = session.NewSessionManager(session.GetSession(r, w))
			// fmt.Fprintf(w, "Session Id: %d", sm.UserLoggedIn())
			handler(w, r, dbSource, sm)
		}
	}

	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Router
	mux.HandleFunc("/", handlerWrapper(model.HomePageHandler))
	mux.HandleFunc("/login", handlerWrapper(model.LoginPageHandler))
	mux.HandleFunc("/register", handlerWrapper(model.RegisterPageHandler))
	mux.HandleFunc("/test", handlerWrapper(model.TestPageHandler))

	// Post
	mux.HandleFunc("/post/register", handlerWrapper(model.RegisterPostHandler))
	mux.HandleFunc("/post/login", handlerWrapper(model.LoginPostHandler))
	mux.HandleFunc("/post/logout", handlerWrapper(model.LogoutPostHandler))

	// Initialize the server
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
