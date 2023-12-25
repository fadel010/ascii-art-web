package main

import (
	"asciiart/asciiart"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	// Register the handler function for the root route
	http.HandleFunc("/", asciiart.HomeHandler)
	http.HandleFunc("/ascii-art", asciiart.ResultHandler)
	http.HandleFunc("/export", asciiart.ExportHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server and handle errors
	log.Println("Server listening on port", port)
	log.Println("Access to the page on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
