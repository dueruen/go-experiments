package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int

func main() {
	// Serve static files from client folder
	http.Handle("/", http.FileServer(http.Dir("client")))

	http.HandleFunc("/sse/dashboard", dashboardHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cashe")
	w.Header().Set("Connection", "keep-alive")

	counter++

	fmt.Fprintf(w, "data: %v\n\n", counter)
	fmt.Printf("data: %v\n\n", counter)
}
