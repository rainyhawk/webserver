package main

// Run a webserver on port 8080
// The server will serve the message "Hello, World!" on the root path "/"

import (
	"fmt"
	"log"
	"net/http"
)

const Port int = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running on port %d\n", Port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe encountered an error: ", err)
		panic(err)
	}
}
