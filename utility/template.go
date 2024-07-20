package utility

import (
	"fmt"
	"html/template"
)

func ParseTemplates(fileNames ...string) (t *template.Template) {
	var filePaths []string
	for _, fileName := range fileNames {
		filePaths = append(filePaths, fmt.Sprintf("templates/%s.gohtml", fileName))
	}
	return template.Must(template.ParseFiles(filePaths...))
}
