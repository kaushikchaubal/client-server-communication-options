package main

import (
	"fmt"
	"net/http"

	"github.com/r3labs/sse"
)

func main() {
	fmt.Printf("Starting sse server")

	server := sse.New()
	server.CreateStream("messages")

	server.Publish("messages", &sse.Event{
		Data: []byte("pong"),
	})

	// Create a new Mux and set the handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", server.HTTPHandler)

	http.ListenAndServe(":3000", mux)
}
