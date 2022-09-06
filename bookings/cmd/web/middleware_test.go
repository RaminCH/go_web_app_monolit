package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSurf(t *testing.T) {

	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}

}

func TestSessionLoad(t *testing.T) {

	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}

}

//go test -v
//go/src/Web_Part1/bookings/cmd/web(master)$ go test -cover
//go test -coverprofile=coverage.out && go  tool cover -html=coverage.out