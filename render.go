package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate is used to render the HTML template pages.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template", err)
		return
	}
}
