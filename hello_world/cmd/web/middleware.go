package main

import (
	"fmt"
	"net/http"
)

// WriteToConsole will write something to console (here - "Hit the page") each time when somebody hits the page
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
