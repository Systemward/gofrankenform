package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Set up a handler function to handle incoming requests
    http.HandleFunc("/", helloWorldHandler)

    // Start the HTTP server on port 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    // Return a "Hello, World!" response
    fmt.Fprintln(w, "Hello, World!")
}
