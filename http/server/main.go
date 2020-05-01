package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}


func main() {
	fmt.Printf("Starting http server")
	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":3000", nil)
}