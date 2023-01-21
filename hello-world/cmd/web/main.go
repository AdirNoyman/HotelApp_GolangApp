package main

import (
	"fmt"
	"hello-world/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s 😎🤟", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
