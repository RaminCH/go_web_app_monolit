package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

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