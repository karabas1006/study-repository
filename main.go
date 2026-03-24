package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	fmt.Println("HTTP server started!")

	err := http_server.StartHTTPServer()
	if err != nil {
		fmt.Println("HTTP server error:", err)
	} else {
		fmt.Println("HTTP stopped.")
	}
}
