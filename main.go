package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main entry point that starts the web server on 8080 port.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting the web server on %v port.", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
