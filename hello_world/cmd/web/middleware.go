package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

// WriteToConsole will write something to console (here - "Hit the page") each time when somebody hits the page
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   //entire site
		Secure:   false, //false cz we are not running https
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
