package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handle the "/" route for GET and POST requests
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// formHandler handles GET requests and serves the form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the HTML form
		http.ServeFile(w, r, "Templates/hh.html")
	}
}

// submitHandler handles POST requests and processes the form data
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		name := r.FormValue("name")
		message := r.FormValue("message")

		// Display the form values
		fmt.Fprintf(w, "Received POST request!\n")
		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Message: %s\n", message)
	}
}
