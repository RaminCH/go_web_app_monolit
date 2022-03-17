package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RaminCH/go-course/pkg/config"
	"github.com/RaminCH/go-course/pkg/handlers"
	"github.com/RaminCH/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false 

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) // "render component" access to "app config"

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run cmd/web/*go
// Starting application on port: :8080


// This is the about page

// This is a test line

// This is a test line2  - is added via "app.UseCache = false"
