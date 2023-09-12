package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is used to render the HTML template pages.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template", err)
		return
	}
}
