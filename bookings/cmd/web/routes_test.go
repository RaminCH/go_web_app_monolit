package main

import (
	"fmt"
	"testing"

	"github.com/RaminCH/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}

//go test -v
//go/src/Web_Part1/bookings/cmd/web(master)$ go test -cover
//go test -coverprofile=coverage.out && go  tool cover -html=coverage.out
