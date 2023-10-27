package main

import (
	"asciiart/asciiart"
	"log"
	"net/http"
)

func main() {
	// Register the handler function for the root route
	http.HandleFunc("/", asciiart.HomePage)
	http.HandleFunc("/asciiart", asciiart.Result)

	// Serve static files (style.css)
	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	log.Println("listening...")
	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
