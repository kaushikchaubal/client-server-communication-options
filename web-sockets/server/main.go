package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func ping(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Read:", err)
			break
		}
		fmt.Printf("Received: %s", message)
		
		err = c.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("Writing:", err)
			break
		}
	}
}

func main() {
	fmt.Printf("Starting web-sockets server")
	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":3000", nil)
}