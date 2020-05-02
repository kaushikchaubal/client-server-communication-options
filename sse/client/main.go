package main

import (
	"fmt"

	"github.com/r3labs/sse"
)

func main() {

	client := sse.NewClient("http://localhost:3000/events")
	// client.Connection.Transport = &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	client.Subscribe("messages", func(msg *sse.Event) {
		// Got some data!
		fmt.Println(string(msg.Data))
	})
}
