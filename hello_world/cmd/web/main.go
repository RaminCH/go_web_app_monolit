package main

import (
	"fmt"
	"net/http"

	"github.com/RaminCH/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run cmd/web/*.go
// Starting application on port: :8080
