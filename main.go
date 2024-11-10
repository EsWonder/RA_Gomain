package main

import (
    "fmt"      // Importing the fmt package for formatted I/O
    "net/http" // Importing the net/http package for handling HTTP requests
)

// handler function is invoked when the "/" route is accessed. It writes "Hola Mundo desde Go!" to the response.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola Mundo desde Go!") // Writes the message to the response
}

// main function sets up the HTTP server and listens for requests on port 8080.
func main() {
    // http.HandleFunc registers the handler function for the root URL ("/").
    http.HandleFunc("/", handler)
    
    // http.ListenAndServe starts the HTTP server on port 8080 and waits for requests.
    http.ListenAndServe(":8080", nil)
}

