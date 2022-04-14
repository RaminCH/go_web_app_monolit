package main

import (
	"net/http"

	"github.com/RaminCH/bookings/pkg/config"
	"github.com/RaminCH/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	//after installing "go get -u github.com/go-chi/chi" - import it from mod via "github.com/go-chi/chi" then
	//delete "pat" via "go mod tidy"

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) //chi middleware(method) that gracefully absorbs panics and prints the stack trace
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//showing assets in images (static folder)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
