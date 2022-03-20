package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RaminCH/go-course/pkg/config"
	"github.com/RaminCH/go-course/pkg/handlers"
	"github.com/RaminCH/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//session
	//change it to true when in production mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// session.Cookie.Secure = false
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//end session

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

	srv := &http.Server{ //srv - serving
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run cmd/web/*.go
// Starting application on port: :8080
