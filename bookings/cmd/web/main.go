package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RaminCH/bookings/internal/config"
	"github.com/RaminCH/bookings/internal/handlers"
	"github.com/RaminCH/bookings/internal/models"
	"github.com/RaminCH/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
	log.Fatal(err)
}
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	//session
	//change it to true when in production mode
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//end session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) // "render component" access to "app config"


	return nil
}

// ramie@ramie:~/go/src/Web_Part1/hello_world(master)$ go run cmd/web/*.go
// Starting application on port: :8080

// visit first about page - it will request to visit home page
// then after visiting home page just go back to about page -> you will see your ip address
