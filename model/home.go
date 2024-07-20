package model

import (
	"auth/database"
	"auth/page"
	"auth/session"
	"auth/utility"
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request, dbSource *database.Source, sm *session.Manager) {
	var pageData *page.Data = page.NewData("Home Page", sm)
	var t *template.Template = utility.ParseTemplates("layout", "home")
	_ = t.ExecuteTemplate(w, "layout", *pageData)
}
