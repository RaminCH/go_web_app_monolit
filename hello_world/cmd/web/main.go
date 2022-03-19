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

	

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	

	srv := &http.Server{		//srv - serving
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

// ramin@ramin:~/go/src/Web_Part1/hello_world(master)$ go run cmd/web/*.go
// Starting application on port: :8080
// Hit the page
// Hit the page							//loaded/refreshed page 3 times (check middleware.go and routes.go)
// Hit the page
// ^Csignal: interrupt


