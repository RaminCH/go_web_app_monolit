package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is: %d", sum))
}

//main is the main application function
func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}


//In terminal
// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run main.go
// Starting application on port: :8080


//in browser
// http://localhost:8080/
// This is the home page

// http://localhost:8080/about
// This is the about page and 2 + 2 is: 4

//Ctrl+C - cancel the process