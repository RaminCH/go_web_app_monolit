package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run main.go
// Starting application on port: :8080

//In browser
// http://localhost:8080/
// This is the home page

// This is the test line

// http://localhost:8080/about
// This is the about page

// This is a test line
