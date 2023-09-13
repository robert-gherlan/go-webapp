package main

import (
	"log"
	"net/http"

	"github.com/robert-gherlan/go-webapp/pkg/config"
	"github.com/robert-gherlan/go-webapp/pkg/handlers"
	"github.com/robert-gherlan/go-webapp/pkg/render"
)

const portNumber = ":8080"

// main is the main entry point that starts the web server on 8080 port.
func main() {
	var app config.AppConfig

	// Init the template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Cannot create the template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	// Init the handlers repo
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting the web server on %v port.", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
